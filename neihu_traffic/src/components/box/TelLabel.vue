<template>
    <div class="mapLabelBox">
        <div>
            <h6>信令推估人口</h6>
            <p>Telecom Estimated Population</p>
            <div class="patio">0%</div>
            <div class="gradient">
                <div 
                    v-for="(item) in speed" 
                    :key="item.text"
                    :style="{
                        'background': `linear-gradient(180deg, #fff, ${item.color})`
                    }"
                >
                    <div>{{item.ch}}</div>
                    <div>{{item.en}}</div>
                </div>
            </div>
            <div class="patio">100%</div>
        </div>
    </div>
</template>

<script>
export default {
    data(){
        return {
            speed: [
                {ch: '工作人口', en: 'Working',  color: '#eaa3a5'},
                {ch: '混合人口', en: 'Mixed',  color: '#8e8cb2'},
                {ch: '居住人口', en: 'Resident',  color: '#adb4ce'},
            ]
        }
    }
}
</script>
<style lang="scss" scoped>
    .mapLabelBox{
        left: 1rem;
        bottom: 1rem;
        z-index: 0;
        >div{
            width: 13rem;
        }
        .patio{
            text-align: right;
            font-size: 0.6rem;
        }
        .gradient{
            display: flex;
            flex-direction: row;
            flex-wrap: wrap;
            width: 90%;
            height: 7rem;
            margin: auto;
            >div{
                opacity: 0.9;
                display: flex;
                flex-direction: column;
                justify-content: flex-end;
                align-items: center;
                flex-basis: 33%;
                font-size: 0.75rem;
            }
        }
    }
</style>

