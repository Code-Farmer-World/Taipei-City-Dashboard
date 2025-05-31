<template>
	<div class="chartLabel">
		<span>Min</span>
		<div
			:style="{
				background: `linear-gradient(90deg, ${gradient[0]}, ${gradient[1]})`
			}"
		>
			{{text}}
		</div>
		<span>Max</span>
	</div>
</template>

<script>
export default {
	props: {
        text: {
            type: String,
			default: ''
        },
        gradient: {
            type: Array,
			default: () => ['#e5eed6', '#208e4f']
        }
    },
}
</script>
<style lang="scss" scoped>
@import '@/assets/scss/variables.scss';
.chartLabel{
    @extend %horizontalSpaceAround;
	align-items: center;
}
</style>

