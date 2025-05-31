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
        },
        gradient: {
            type: Array,
			default: () => ["#999", "#111"]
        }
    },
}
</script>
<style lang="scss" scoped>
@import '@/assets/scss/variables.scss';
.chartLabel{
    @extend %horizontalSpaceAround;
	align-items: center;
	>div{
		position: relative;
		border-radius: 0.7rem;
		&:before,
		&:after{
			content: '';
			position: absolute;
			top: 0.1rem;
			border: 1px solid #888;
			width: 0.7rem;
			height: 0.7rem;
			border-radius: 50%;
		}
		&:before{
			left: 0;
		}
		&:after{
			right: 0;
		}
	} 
}
</style>

