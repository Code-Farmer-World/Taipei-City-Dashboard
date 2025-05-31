<template>
    <div class="apexChartContainer columnBus">
        <apexchart 
            type="bar"
            height="250"
            :options="chartOptions" 
            :series="seriesData"
        />
    </div>
</template>

<script>
import { TranData, ParseRatio } from '@/assets/js/transport.js'
// const ClearData = TranData.sort((a, b)=> b['bus_avg'] - a['bus_avg']).slice(0, 15)
const AvgSum = 8750
const ClearData = TranData.slice(0,15).map(item => {
    return {
        ...item,
        ratio: ParseRatio(item['bus_avg'], AvgSum, 0)
    }
})
export default {
    data(){
        return {
            seriesData: [{
                name: this.$t("charts.columnBasicBus.name"),
                data: ClearData.map(item => item['ratio'])
            }],
            chartOptions: {
                colors: ['#5b955b'],
                title: {
                    text: this.$t("charts.columnBasicBus.title"),
                    margin: 0,
                    style: {
                        fontSize: '12px'
                    }
                },
                subtitle: {
                    text: `${this.$t("count")}:${AvgSum} ${this.$t("per")}`,
                    margin: 0,
                    offsetY: 15
                },
                plotOptions: {
                    bar: {
                        horizontal: false,
                        columnWidth: '60%',
                    },
                },
                xaxis: {
                    categories: ClearData.map(item => item['name']),
                    labels: {
                        rotate: -90,
                        style: {
                            fontSize: '10px',
                        },
                        minHeight: 120,
                        offsetY: -5
                    }
                },
                yaxis: {
                    labels: {
                        maxWidth: 20,
                        style: {
                            fontSize: '8px',
                        },
                        formatter: (val) => { return `${val}%` }
                    }
                },
                dataLabels: {
                    enabled: false
                },
                tooltip:{
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
.apexChartContainer.columnBus{
    width: 100%;
    height: 190px;
    overflow: hidden;
}
</style>
