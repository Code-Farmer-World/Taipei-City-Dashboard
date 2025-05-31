<template>
    <div class="apexChartContainer donut">
        <apexchart 
            id="donutChart" 
            type= "donut"
            height="200" 
            :options="chartOptions" 
            :series="chartData"
        />
    </div>
</template>

<script>
export default {
    props: {
        chartData: {
            type: Array,
            default: () => []
        }
    },
    data(){
        return {
            chartOptions: {
                colors: ['#6b9a80','#72be95', '#8bc08b', '#f44646'],
                labels: this.$i18n.locale === 'zh-TW'?['公車', '捷運', 'Youbike','私人運具']: ['Bus', 'MRT', 'Youbike', 'Private'],
                legend: {
                    show: false,
                    position: 'right',
                    formatter:(seriesName, opts) =>{
                        const chartVal = opts.w.globals.seriesPercent[opts.seriesIndex][0]
                        return `${seriesName}: ${chartVal? parseFloat(chartVal.toFixed(1)): 0}%`
                    }
                },
                dataLabels: {
                    enabled: true,
                    formatter: (val, opts) => {
                        if(val <= 5)return
                        const label = this.chartOptions.labels[opts.seriesIndex]
                        const Ratio = `${parseFloat(val.toFixed(1))}%`
                        return `${label} ${Ratio}`
                    },
                    style: {
                        colors: ['#fff'],
                        fontWeight: 'normal',
                    },
                    background: {
                        enabled: true,
                        foreColor: '#000',
                        borderColor: '#666',
                        borderWidth: 1
                    }
                },
                tooltip:{
                    enabled: true
                },
                responsive: [{
                    breakpoint: 380,
                    options: {
                        legend: {
                            show: false
                        }
                    }
                }]
            }
        }
    },
    watch:{
		chartData(newVal, oldVal){
            if(newVal){
                this.updateChart()
            }
        }
    },
    methods:{
        updateChart(){
            if(this.chartData.length === 0) return
            if(!this.chartOptions.legend.show){
                this.chartOptions = {
                    ...this.chartOptions,
                    legend: {
                        ...this.chartOptions.legend,
                        show: true
                    }
                }
            }
        }
    },
    created(){
        this.updateChart()
    }
}
</script>
<style lang="scss" scoped>
@import '@/assets/scss/variables.scss';
.apexChartContainer{
    &.donut{
        width: $mapAsideWidth;
        .apexcharts-legend-series{
            height: auto;
            line-height: .8rem;
            margin-bottom: 0 !important;
        }
    }
}
</style>
