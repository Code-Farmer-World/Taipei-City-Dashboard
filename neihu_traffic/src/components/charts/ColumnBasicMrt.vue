<template>
    <div class="apexChartContainer columnMrt">
        <apexchart 
            type="bar"
            height="400"
            :options="chartOptions" 
            :series="seriesData"
        />
    </div>
</template>

<script>
import { MrtData, MrtColor, ParseRatio } from '@/assets/js/transport.js'
const WenhuData = []
const NonWenhuData = []
const DanshuiData = []
const AvgSum = 18161
MrtData.map(item => {
    if(item.type === '文湖線'){
        WenhuData.push({
            ...item,
            ratio: ParseRatio(item.avg, AvgSum),
            color: MrtColor[item.type]
        })
    }else{
        NonWenhuData.push({
            ...item,
            ratio: ParseRatio(item.avg, AvgSum),
            color: MrtColor[item.type]
        })
    }
    if(item.type === '淡水信義線'){
        DanshuiData.push({
            ...item,
            ratio: ParseRatio(item.avg, AvgSum),
            color: MrtColor[item.type]
        })
    }
})
export default {
    props: {
        article: {
            type: Number,
            default: 1
        }
    },
    data(){
        return {
            sourceData: [],
            chartOptions: {},
            seriesData: [{
                name: '捷運轉乘佔比',
                data: []
            }],
            title: '',
            categories: [],
            yaxisColors: [],
            zone : [],
            subtitle: ''
        }
    },
    mounted(){
        this.updateData()
    },
    watch:{
		article(newVal, oldVal){
            this.updateData()
        }
    },
    methods:{
        updateData(){
            switch (this.article) {
                case 3:
                    this.title = this.$t("charts.columnBasicMrt.title3")
                    this.subtitle = `${this.$t("accountAll")}32.5%`
                    this.sourceData = NonWenhuData.slice(0, 20)
                    break;
                case 4:
                    this.title = this.$t("charts.columnBasicMrt.title4")
                    this.subtitle = `${this.$t("accountAll")}2.8%`
                    this.sourceData = DanshuiData.slice(0, 20)
                    break;
                default:
                    this.title = this.$t("charts.columnBasicMrt.title2")
                    this.subtitle = `${this.$t("accountAll")}45.95%`
                    this.sourceData = WenhuData.slice(0, 20)
                    break;
            }

            this.seriesData[0]['data'] = []
            this.categories = []
            this.yaxisColors = []
            this.sourceData.map(item => {
                this.seriesData[0]['data'].push(item['ratio'])
                this.categories.push(item['name'])
                this.yaxisColors.push(item['color'])
            })
            this.updateChart()
        },
        updateChart(){
            this.chartOptions = {
                plotOptions: {
                    bar: {
                        horizontal: true,
                        dataLabels: {
                            position: 'bottom'
                        }
                    }
                },
                colors: ['#ddd'],
                title: {
                    text: this.title,
                    margin: 0,                    
                    style: {
                        fontSize: '12px'
                    }
                },
                subtitle: {
                    text: this.subtitle,
                    margin: 0,
                    offsetY: 15
                },
                xaxis: {
                    categories: this.categories,
                    labels: {
                        minHeight: 90,
                        style: {
                            fontSize: '11px',
                        }
                    },
                    title: {
                        text: `${this.$t("accountFor")}(％)`,
                        style: {
                            fontSize: '10px',
                        }
                    }
                },
                yaxis: {
                    labels: {
                        maxWidth: 90,
                        style: {
                            fontSize: '11px',
                            colors: this.yaxisColors
                        }
                    }
                },
                dataLabels: {
                    enabled: true,
                    style: {
                        fontSize: '11px',
                        fontWeight: 'normal',
                        colors: ['#555']
                    },
                    textAnchor: 'start',
                    formatter: (val, opt) => {
                        const Target = this.sourceData[opt.dataPointIndex]
                        if(this.article === 3){
                            return Target && Target.zone ? Target.zone: Target.mrt_type
                        }else{
                            return Target && Target.mrt_type ? `${Target.mrt_type}`: ''
                        }
                    }
                },
                tooltip:{
                    marker: {
                        show: false
                    },
                    fixed: {
                        enabled: true,
                        position: 'bottomRight',
                        offsetX: 0,
                        offsetY: -50,
                    },
                    x: {
                        formatter: (val, { series, seriesIndex, dataPointIndex, w }) => {
                            const Target = this.sourceData[dataPointIndex]
                            if(!(Target && Target.type))return val
                            return `${val}(${Target.type? Target.type: ''}${Target.mrt_type? '、'+Target.mrt_type: ''} )`
                        }
                    },
                    y: {
                        formatter: (val) => (`${val}%`)
                    }
                }
            }
        } 
    }
}
</script>

<style lang="scss"> 
.apexChartContainer.columnMrt{
    width: 100%;
    overflow: hidden;
    .apexcharts-bar-area:hover {
        fill: #999;
    }
}
</style>
