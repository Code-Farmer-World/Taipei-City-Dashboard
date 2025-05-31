<template>
    <div class="loadingContainer" :class="{loadStart}">
        <div class="loadingBox">
            <div/><div/><div/>
        </div>
    </div>
</template>
<script>
export default {
    props: {
        loadStart: {
            type: Boolean,
            default: false
        }
    }
}
</script>
<style lang="scss">
@import '@/assets/scss/variables.scss';
    .loadingContainer{
        @extend %absoluteLeftLayout;
        width: 100%;
        height: 100%;
        z-index: 10;
        @extend %verticalCenter;

        background-color: rgba($blackColor, 0.2);
        transition-property: opacity, display;
        transition-duration: 3s, 1s;
        animation-iteration-count: infinite, infinite;
        animation-direction: alternate, alternate;

        display: none;
        opacity: 0;
        &.loadStart{
            display: flex;
            opacity: 1;
        }
        .loadingBox{
            @extend %horizontalCenter;
            align-items: center;
            height: 100%;
            >div {
                margin: 10px;
                width: 50px;
                height: 50px;
                border-radius: 50%;
                background-color: $whiteColor;
                animation-duration: 0.8s;
                animation-iteration-count: infinite;
                &:nth-child(1){
                    animation-name: load-one;
                }
                &:nth-child(2){
                    animation-name: load-two;
                }
                &:nth-child(3){
                    animation-name: load-three;
                }
            }
        }
    }
    @keyframes load-one {
        30% {
            transform: translateY(-50%);
        }
    }
    @keyframes load-two {
        50% {
            transform: translateY(-50%);
        }
    }
    @keyframes load-three {
        70% {
            transform: translateY(-50%);
        }
    }
</style>