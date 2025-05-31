<template>
    <div class="apexChartContainer">
        <apexchart :options="chartOptions" :series="series"/>
    </div>
</template>

<script>
const seriesData = [{
    name: 'Net Profit',
    data: [44, 55, 57, 56, 61, 58, 63, 60, 66]
}, {
    name: 'Revenue',
    data: [76, 85, 101, 98, 87, 105, 91, 114, 94]
}]
export default {
    props: {
        load: {
            type: Boolean,
            default: false
        }
    },
    data(){
        return {    
            series: [],
            chartOptions: {
                chart: {
                    type: 'bar',
                    height: 350
                },
                plotOptions: {
                    bar: {
                        horizontal: false,
                        columnWidth: '55%',
                        endingShape: 'rounded'
                    },
                },
                dataLabels: {
                    enabled: false
                },
                stroke: {
                    show: true,
                    width: 2,
                    colors: ['transparent']
                },
                xaxis: {
                    categories: ['Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct']
                },
                yaxis: {
                    title: {
                        text: '$ (thousands)'
                    }
                },
                fill: {
                    opacity: 1
                },
                tooltip: {
                    y: {
                        formatter: function (val) {
                            return "$ " + val + " thousands"
                        }
                    }
                }
            }
        }
    },
    watch:{
		load(newVal, oldVal){
            if(newVal){
                this.series = seriesData
            }else{
                this.series = []
            }
        }
    }
}
</script>
