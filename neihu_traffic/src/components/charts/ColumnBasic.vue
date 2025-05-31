<template>
    <div class="chartActionBox">
        <div class="toggleBtn" :data-attr="$t('sort')">
            <button :class="{active:sortType == 0}" @click="sortTypeToggle(0)">{{$t("charts.columnBasic.sort0")}}</button>
            <button :class="{active:sortType == 1}" @click="sortTypeToggle(1)">{{$t("charts.columnBasic.sort1")}}</button>
        </div>
        <div class="toggleBtn" :data-attr="$t('neihu')">
            <button :class="{active:neiHuShow}" @click="neiHuShowToggle(true)">{{$t("show")}}</button>
            <button :class="{active:!neiHuShow}" @click="neiHuShowToggle(false)">{{$t("hide")}}</button>
        </div>
        <div class="toggleBtn" :data-attr="$t('chart')">
            <button :class="{active:!stackedChart}" @click="stackedChartToggle(false)">{{$t("charts.columnBasic.stacked0")}}</button>
            <button :class="{active:stackedChart}" @click="stackedChartToggle(true)">{{$t("charts.columnBasic.stacked1")}}</button>
        </div>
    </div>
    <div v-if="chartReload" class="apexChartContainer basic">
        <apexchart 
            type='bar'
            height='100%'
            :options="chartOptions" 
            :series="chartValue"/>
    </div>
</template>

<script>
import {PrivateT, PublicT} from '@/assets/js/transport.js'
export default {
    data(){
        return {
            chartReload: false,
            seriesType: {
                '0':{index:'private_transport',name: this.$t("charts.columnBasic.sort0")},
                '1':{index:'public_transport',name: this.$t("charts.columnBasic.sort1")}
            },
            sortType: '0',  
            neiHuShow: true,
            stackedChart: false,
            transData: [],

            categories: [],
            seriesData: [],

            chartValue: []
        }
    },
    props: {
        load: {
            type: Boolean,
            default: false
        }
    },
    computed:{
        chartOptions(){
            const Option = {
                colors: ['#fb6b6d', '#73ba8c'],
                chart:{
                    stacked: (this.stackedChart)? true: false,
                    stackType: (this.stackedChart)? '100%': 'normal'
                },
                plotOptions: {
                    bar: {
                        horizontal: false
                    }
                },
                xaxis: {
                    categories: this.categories,
                    labels: {
                        rotate: 0,
                        hideOverlappingLabels: false
                    }
                },
                yaxis: {
                    max: (e) => e
                }
            }
            return {
                ...Option,
                dataLabels: {
                    enabled: true,
                    style: {
                        fontSize: '10px',
                        fontWeight: 'normal',
                        colors: ['#666']
                    },
                    offsetY: -20,
                    formatter: (val, opts) => {
                        return (val%1 === 0)? val: `${parseFloat(val.toFixed(1))}%`
                    },
                },
                legend: {
                    position: 'top',
                    horizontalAlign: 'left'
                }
            }
        }
    },
    methods:{
        sortTypeToggle(val){
            this.sortType = val
            this.updateChart()
        },
        neiHuShowToggle(boolen){
            this.neiHuShow = boolen
            this.updateChart()
        },
        stackedChartToggle(boolen){
            this.stackedChart = boolen
        },
        updateChart(){
            this.chartReload = false
            this.chartValue = []
            this.seriesData = []
            //依據大眾運輸｜私人運具排序
            // const targetKey = this.seriesType[this.sortType]['index']
            // const clearData = transportData.sort((a, b)=>{
                //     return b[targetKey] - a[targetKey]    //從大到小
            //     // return a[targetKey] - b[targetKey] //從小到大
            // })
            // 依據上述算法資料已排序，排出兩者皆為0的鄉鎮
            const clearData = this.sortType == 0? PrivateT: PublicT
            //是否顯示內湖資料：內湖都在第一筆
            this.transData = this.neiHuShow? clearData.slice(0, 33): clearData.slice(1, 34)
            
            // 依據大眾運輸｜私人運具顯示
            Object.keys(this.seriesType).map(key=>{
                const {name, index} = this.seriesType[key]
                this.seriesData.push({
                    name: name,
                    data: this.transData.map(item => item[index])
                })
            })
            this.categories = this.transData.map(item => {
                return this.$i18n.locale === 'zh-TW'? item['TOWNNAME']: item['TOWNENAME']
            })
            this.chartValue = this.seriesData
            this.chartReload = true

            this.$gtag.event('update_column_chart', {
                event_category: 'neihu_traffic_action'
            })
        }
    },
    mounted(){
        this.updateChart()
    }
}
</script>

<style lang="scss" scoped>
@import '@/assets/scss/variables.scss';
.apexChartContainer{
    display: block;
    height: 32.5rem;
    @include scrollbar_style;
    width: calc(100vw - 4rem);
    overflow-x: auto;
    overflow-y: hidden;
    @media screen and (max-width:501px){  // For mobile
        height: 50vh;
    }
}
</style>
<style lang="scss"> 
.apexChartContainer.basic .vue-apexcharts{ 
    width: 200rem;
    height: 100%;
    .apexcharts-legend{
        overflow: hidden;
    }
    .apexcharts-toolbar{
        right: -1rem !important;
    }
}
</style>
