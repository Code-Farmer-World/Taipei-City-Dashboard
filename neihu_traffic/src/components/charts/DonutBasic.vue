<template>
    <div class="apexChartContainer donut">
        <apexchart 
            type= "donut"
            height="200" 
            :options="chartOptions" 
            :series="series"
        />
    </div>
</template>

<script>
const chartDatas = [59, 16, 25]
const chartColors = ['rgb(203, 15, 15)','rgb(229, 135, 135)', 'rgb(238, 180, 180)']
export default {
    data(){
        return {
            series: chartDatas,
            chartOptions: {
                labels: this.$t("charts.donutBasic.label"),
                colors: chartColors,
                legend: {
                    show: false
                },
                dataLabels: {
                    enabled: true,
                    formatter: (val) => {
                        const dataIndex = chartDatas.findIndex(item => item === val)
                        return `${this.$t("charts.donutBasic.label")[dataIndex]} ${val}%`
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
                    enabled: false
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
