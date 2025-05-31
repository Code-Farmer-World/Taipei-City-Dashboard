<template>
    <div class="auther" data-aos="fade-up" data-aos-duration="800" data-aos-delay="400">
        <h6>{{$t('footer.participate')}}</h6>
        <p v-html="$t('footer.menber')"/>
        <h6>{{$t('footer.collaborate')}}</h6>
        <p v-html="$t('footer.unit')"/>
    </div>
</template>

<script>
export default {
    data(){
        return {}
    }
}
</script>
<style lang="scss" scoped>
.auther{
    padding: 2rem;
    h6{
        color: #7f7f7f;
    }
}
</style>