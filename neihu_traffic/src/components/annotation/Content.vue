<template>
    <div :class="[
        'annotationDropdown',
        fix,
        {open: dropdown},
    ]" >
        <div class="annotationContent" >
            <p v-html="content"/>
        </div>
        <img 
            v-if="contentImg !== ''"
            class="annotationImage"
            :src="require(`@/assets/img/${contentImg}`)" :alt="text"
        />
    </div>
</template>
<script>
const includeEl = ["annotationBtn", "annotationDropdown", "annotationContent", "annotationImage"]
export default {
	props:{
        dropdown: {
            type: Boolean,
            default: false
        },
        fix: {
            type: String,
            default: ''
        },
		content: {
            type: String,
            default: ''
        },
        contentImg: {
            type: String,
            default: ''
        },
		text: {
            type: String,
            default: ''
        }
	},
    mounted() {
        document.addEventListener("click", this.bodyClosePopover)
    },
    beforeUnmount() {
        document.removeEventListener("click", this.bodyClosePopover)
    },
    methods: {
        bodyClosePopover(e) {
            if(!this.dropdown)return
            if(e.target && e.target.classList.value && includeEl.includes(e.target.classList.value)) return
            this.$emit('update', false)
        }
    }
}
</script>
<style lang="scss" scoped>
@import '@/assets/scss/annotation.scss';
</style>