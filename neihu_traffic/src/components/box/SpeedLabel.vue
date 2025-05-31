<template>
    <div class="mapLabelBox">
        <div>
            <h6>路段平均時速</h6>
            <p>Avg Speed km/hr</p>
            <div class="legendBox" v-for="(item) in showSpeed" :key="item.text">
                <span :style="{
                    'background-color': item.color
                }"/>
                <span>{{item.text}}km/hr</span>
            </div>
            <div class="toggleBtn" :data-attr="$t('box.speedLabel.roadSpeed')">
                <button :class="{active:!lowSpeed}" @click="updateLowSpeed(false)">{{$t('all')}}</button>
                <button :class="{active:lowSpeed}" @click="updateLowSpeed(true)">{{$t('box.speedLabel.filterRow')}}</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data(){
        return {
            speed: [
                {text: '0-30', color: '#941c20'},
                {text: '30-40', color: '#d23d0d'},
                {text: '40-60', color: '#aa791c'},
                {text: '>60', color: '#53a755'},
            ],
            showSpeed: []
        }
    },
    props:['lowSpeed'],
    watch:{
        lowSpeed: {
            immediate: true,
            handler(newVal) {
                this.showSpeed = newVal? [this.speed[0]]: this.speed      
            }
        }
    },
    methods:{
        updateLowSpeed(boolen){
            this.$emit('update', boolen)
            this.$gtag.event('update_speed', {
                event_category: 'neihu_traffic_action'
            })
        }
    }
}
</script>
<style lang="scss" scoped>
    .mapLabelBox{
        left: 1rem;
        bottom: 1rem;
        z-index: 0;
        .legendBox{
            height: 1rem;
            span{
                &:first-child{
                    width: 3rem;
                    height: 0.2rem;
                }
            }
        }
        .toggleBtn{
            margin-left: 1rem;
            margin-top: 2rem;
        }
    }
</style>

