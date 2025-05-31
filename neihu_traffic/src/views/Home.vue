<template>
	<main ref="home_container">
		<article 
			id="main_scrollama" 
			ref="scrollama_container"
			:class="{langEn: !langZh}"
		>
			
			
			
			<div data-step="7" class="step mapContainer">
				<StepMap/>
			</div>
		
		
		</article>
		<AsideBox :container-height="containerHeight"/>
	</main>
</template>

<script>
import { defineAsyncComponent } from 'vue'
// scrollama
import "intersection-observer"
import scrollama from "scrollama"

import AOS from "aos"
import "aos/dist/aos.css"

export default {
	name: "HomePage",
	components:{
		ImgContainer: defineAsyncComponent(() => import('@/components/content/ImgContainer.vue')),
		StepMap: defineAsyncComponent(() => import('@/components/content/StepMap.vue'))
	},
	data() {
		return {
			opts: {}
		}
	},
	computed: {
		langZh(){
            return this.$i18n.locale === 'zh-TW'
        },
        mobileDevice(){
            return this.$store.state.mobileDevice
        },
		currStep() {
			return this.$store.state.step
		},
		currStepProgress() {
			return this.$store.state.progres
		},
		containerHeight(){
			return this.$refs.home_container? this.$refs.home_container.offsetHeight: 0
		},
		imgContainerFixed(){
			if(this.currStep == 0)return false
			if(this.currStep > 4)return false
			if(this.mobileDevice)return true
			return this.$store.state.contentEnter
		}
	},
  	created() {
    	AOS.init({})
    },
	mounted() {
		const childNodes = this.$refs.scrollama_container.childNodes
		if(typeof childNodes === 'object'){
			const step = Object.values(childNodes).filter(item => {
				return typeof item === 'object' && item.hasAttribute('data-step')
			})
			this.opts = Object.assign({},  {
				step: '.step',
				progress: true
			}, this.$attrs)
		}
		this.setupScroller()
		window.addEventListener('resize', this.handleResize)
	},
	beforeUnmount() {
		this.$store.commit('updateStep', 0)
		this._scroller.destroy()
		window.removeEventListener('resize', this.handleResize)
	},
	methods: {
		setupScroller() {
			this._scroller = scrollama()
			this._scroller
			.setup(this.opts)
			.onStepExit(({ element, index }) => {
				if(!this.mobileDevice)return
				if(index == 0) this.$store.commit('updateStep', 0)
			})
			.onStepProgress(({element, progress}) => {
				this.$store.commit('updateStep', element.dataset.step)
				this.$store.commit('updateProgres', progress)
			})
		},
		handleResize () {
			this._scroller.resize()
		}
	}
}
</script>

<style lang="scss">
@import '@/assets/scss/home.scss';
</style>
