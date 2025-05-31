<template>
	<div class="carousel-header">
		<div class="carousel-btn">
			<button
				v-for="(slide) in tab" 
				:key="slide.index"
				@click="selectPage(slide.index)"
				:class="{
					'slideActive': (currentTab === slide.index)
				}"
			>
				{{slide.name}}
			</button>
		</div>
	</div>
</template>

<script>
export default {
    props: {
        tab: {
            type: Array
        },
		currentTab: {
			type: String
		}
    },
	methods: {
		selectPage(val){
			this.$emit('update', val)
			this.$gtag.event('update_carousel', {
                event_category: 'neihu_traffic_action'
            })
		}
	}
}
</script>