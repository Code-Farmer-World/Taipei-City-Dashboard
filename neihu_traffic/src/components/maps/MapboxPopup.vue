<template>
    <div class="popupBox">
        <div v-for="(feature, i) in featuresData" :key="i" class="featuresDataBox">
            <div v-for="(featureItem, featureKey) in feature" :key="featureKey">
                <span>{{featureKey}}</span>
                <b>{{featureItem}}</b>
            </div>
        </div>
    </div>
</template>

<style lang="scss" >
@import '@/assets/scss/variables.scss';
.mapboxgl-popup-content {
    width: 13rem;
    font-size: 0.9rem;
    padding: 1rem 0 0 .5rem;
    button{
        width: 2rem;
        font-size: 1.5rem;
    }
}
.popupHeader{
    color: #6d8a8e;
    text-align: left;
}
.popupBox{
    @include scrollbar_style;
    white-space:nowrap;
    overflow: scroll;
    text-align: left;
}
.featuresDataBox{
    display: inline-block;
    word-wrap:break-word;
    width: 12rem;
    >div{
        margin-bottom: .25rem;
    }
    span{
        display: inline-block;
        width: 8rem;
        padding-right: .5rem;
    }
}
</style>