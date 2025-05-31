<template>
    <div class="apexChartContainer radialsBasic">
        <apexchart 
            type= "radialBar"
            height="280" 
            :options="chartOptions" 
            :series="series"
        />
    </div>
</template>

<script>
const chartDatas = [62, 32, 6]
const chartColors = ['#73ba8c','#a0cfb2', '#c1ddca']
export default {
    data(){
        return {
            series: chartDatas,
            chartOptions: {
                labels: this.$i18n.locale === 'zh-TW'?['公車', '捷運', 'Youbike']: ['Bus', 'MRT', 'Youbike'],
                colors: chartColors,
                plotOptions: {
                    radialBar: {
                        inverseOrder: true,
                        startAngle: 0,
                        endAngle: 270,
                        hollow: {
                            margin: 10,
                            size: '42.5%',
                            background: 'transparent',
                            image: undefined,
                        },
                        dataLabels: {
                            name: {
                                fontSize: '14px',
                                offsetY: -15
                            },
                            value: {
                                fontSize: '18px',
                                offsetY: -5
                            },
                            total: {
                                show: true,
                                fontSize: '12px',
                                color: '#373d3f',
                                label: this.$i18n.locale === 'zh-TW'?'搭乘大眾運輸人口':'Public transport',
                                formatter: (w) => {
                                    return `42,500${this.$i18n.locale === 'zh-TW'?"人": "pl"}`
                                }
                            }
                        }
                    }
                },
                legend: {
                    show: true,
                    inverseOrder: true,
                    height: 100,
                    floating: true,
                    position: 'left',
                    horizontalAlign: 'left',
                    offsetX: 0,
                    offsetY: 0,
                    labels: {
                        useSeriesColors: true,
                    },
                    markers: {
                        width: 5,
                        height: 5,
                        radius: 5
                    },
                    formatter: function(seriesName, opts) {
                        return seriesName + ":  " + opts.w.globals.series[opts.seriesIndex] + "%"
                    }
                },
                itemMargin: {
                    horizontal: 0,
                    vertical: 0
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
    }
}
</script>

<style lang="scss">
@import '@/assets/scss/variables.scss';
.apexChartContainer{
    &.radialsBasic{
        .apexcharts-legend-series{
            height: auto;
            line-height: .8rem;
            margin-bottom: 0 !important;
        }
    }
}
</style>
