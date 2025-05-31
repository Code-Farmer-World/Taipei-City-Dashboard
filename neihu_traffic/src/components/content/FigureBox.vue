<template>
    <figure id="main_figure" :class="currStep > 1? 'fixed': 'absolute'">
        <div class="imgBox">
            <div class="bgBlock bg1 active"
                :class="{
                    activeOpacity: eyecatch
                }"
            />
            <div class="bgBlock"
                :class="{
                    active: textShow2 || continueNext,
                    activeOpacity: eyecatch || continueNext
                }"
                :style="{
                    backgroundImage: `url(${lowSpeed? MapImg2: MapImg3})`
                }"
            />
        </div>
    </figure>
</template>
<script>
export default {
    components:{
	},
    computed: {
        mobileDevice(){
            return this.$store.state.mobileDevice
        },
        currStep() {
			return this.$store.state.step
		},
		currStepProgress() {
			return this.$store.state.progres
		},
    }
}
</script>