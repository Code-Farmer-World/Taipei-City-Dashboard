<template>
	<div class="carousel-image">
		<div 
			v-for="(slide) in tab" 
			:key="slide.index"
			:style="slideSelect(slide.index)"
			:class="{
				'slideActive': (currentTab === slide.index)
			}"
		>
			<div 
				class="imgBox"
				:style="{
					'background-image': 'url('+ require(`@/assets/img/${imgURL}/${slide.index}.jpg`)+')',
				}"
			>
				<template v-if="imgURL === 'tab'">
					<div v-if="currentTab === 'All'" class="imgLabelContainer all"/>
					<div v-else class="imgLabelContainer"
						:style="{
							'background-image': 'url('+ require(`@/assets/img/${imgURL}/mini/${currentTab}.jpg`)+')',
						}"
					/>
				</template>
			</div>
		</div>
		<ChartLabel :text="label" :gradient="gradient"/>
	</div>
</template>

<script>
import { defineAsyncComponent } from 'vue'
export default {
	components:{
		ChartLabel: defineAsyncComponent(() => import('@/components/content/ChartLabel.vue'))
	},
    props: {
        tab: {
            type: Array
        },
		currentTab: {
			type: String
		},
		gradient: {
			type: Array
		},
		label: {
			type: String
		},
		imgURL: {
			type: String
		},
		imgLabel: {
			type: Boolean
		}
    },
	methods: {
		slideSelect(index) {
			const selectIndex = this.tab.findIndex( item => item.index == this.currentTab)
			return {
				left : `${(index-selectIndex)*100}%`
			}
		}
	}
}
</script>
<style lang="scss" scoped>
.imgLabelContainer{
    &.all{ background-image: url('../../assets/img/tab/mini/zone.jpeg'); }
}
</style>
