<template>
	<div class="chartLabel">
		<div class="icon"
			:style="{
				borderWidth: `${borderWidth}px`,
				borderStyle: borderStyle,
				borderColor: borderColor,
			}"
		/>
		<span>{{text}}</span>
	</div>
</template>

<script>
export default {
	props: {
        text: {
            type: String,
        },
        borderWidth: {
            type: Number,
			default: () => 1
        },
        borderColor: {
            type: String,
			default: () => '#ddd'
        },
        borderStyle: {
            type: String,
			default: () => 'solid'
        }
    },
}
</script>
<style lang="scss" scoped>
@import '@/assets/scss/variables.scss';
.chartLabel{
    @extend %horizontalSpaceAround;
	.icon{
		border-radius: 50%;
		width: 1.25rem;
		height: 1.25rem;
		margin: 0 .5rem;
		padding: 0;
	}
}
</style>

