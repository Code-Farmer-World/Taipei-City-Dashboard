<template>
    <div class="mapLabelBox">
        <div>
            <h6>大眾運通勤率</h6>
            <p>Public Transit Commute Rates</p>
            <div class="legendBox" v-for="(item) in rates" :key="item.text">
                <span :style="{
                    'background-color': item.color
                }"/>
                <span>{{item.text}}%</span>
            </div>
            <p>{{$t("charts.columnBasicMrt.name")}}</p>
        </div>
    </div>
</template>

<script>
export default {
    data(){
        return {
            rates: [
                {text: '0-5', color: '#fb6b6d'},
                {text: '6-10', color: '#ed8281'},
                {text: '11-20', color: '#e98f90'},
                {text: '21-30', color: '#ecbabb'},
                {text: '31-50', color: '#d5e4db'},
                {text: '51-70', color: '#c1ddca'},
                {text: '71-100', color: '#a0cfb2'},
                {text: '>100', color: '#73ba8c'},
            ]
        }
    }
}
</script>
<style lang="scss" scoped>
    .mapLabelBox{
        left: auto;
        right: .25rem;
        bottom: .75rem;
        z-index: 1;
        >div{
            width: 14rem;
            height: 18rem;
            > p{
                position: relative;
                width: 100%;
                line-height: 1.6;
            }
        }
        .legendBox{
            width: 100%;
            span{
                line-height: 1;
                &:first-child{
                    width: 1rem;
                    height: 1rem;
                }
            }
        }
    }
</style>

