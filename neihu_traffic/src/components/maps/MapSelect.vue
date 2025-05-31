<template>
    <header class="svgHeader">
        <svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="200 25 200 125" style="enable-background:new 0 0 400 535;" xml:space="preserve">
            <g v-for="item in city" :key="item.index" @click="clickLocation(item)" :class="{active: location && location.index === item.index}">
                <desc>{{item.name}}</desc>
                <path :data-name="item.index" :d="item.path"/>
            </g>
        </svg>
        <ul>
            <li @click="clickLocation(null)" :class="{active: Object.keys(location).length === 0}">
                {{$t('nTw')}}
            </li>
            <br/>
            <li v-for="item in city" :key="item.index" @click="clickLocation(item)" :class="{active: location && location.index === item.index}">
                {{$i18n.locale == 'zh-TW'? item.name: item.en}}
            </li>
        </ul>
    </header>
</template>

<script>
export default {
    props: {
        location: {
            type: Object
        } 
    },
	data() {
		return {
            city: [
                {
                    index: 'taipei_city',
                    countycode: 63000,
                    name: '臺北市',
                    en: 'Taipei',
                    path: 'M317.1,71.5c1.2,4.1-2.5,2.9-4.1,3.9c-1.7,1,2.3,5.2,0,5.8c-7.7,0.4-8.1-3.1-11-5c-2.9-1.9-4.3-3.9-4.3-7 c0-3-1.4-5.5-5.6-8.4c-4.1-2.9-2.9-5.2-0.8-5.6c2.1-0.4,15-8.2,16.4-8.7c1.4-0.5,2.7,1.9,1.9,3.7c-0.8,1.9,0.2,3.9,1,7.7 c0.8,3.7,4.4,2.9,4.3,4.1C314.8,63.4,315.9,67.4,317.1,71.5z'
                },
                {
                    index: 'new_taipei_city',
                    countycode: 65000,
                    name: '新北市',
                    en: 'New Taipei',
                    path: 'M361.5,72.3c-6.8-0.3-9,0.3-9.3-6.2c1.2-11.2-6.2-9.3-13.7-9.3c-2.4,0-4.5-0.4-6.4-1c1.2,2.3,0.3,3.6,0.7,6.3 c0.4,2.9,0.6,4.4,0,6.4c-0.6,2.1-11.8-0.6-15.3-6c-3.5-5.4-3.1-7.9-1.7-8.7c1.1-0.7,4.2-1.7,7.7-3.2c-0.8-1.8-0.2-4.7-3.1-4.4 c-3.4,0.3-5.3-2.5-6.2-8.1c-0.9-5.6-15.8-3.1-18.3-2.2c-2.5,0.9-13,13.4-15.8,15.6c0,0-2.2,4.6-3.7,6.1c-1,1-4.1,1.5-7.5,1.6 c4.1,3.4,5.5,4.9,8.2,5.9c3.3,1.2,5.6,0.6,6.6,2.9c1,2.3,2.7,5.4,0.8,7.3c-1.9,1.9-5.8,4.3-7.7,3.9c-1.9-0.4-2.1,7.9-1.3,10.4 c0.7,2.5,3.8,2.5,2.6,4.6c-1.2,2.1-0.4,5,3.1,4.8c3.5-0.2,5.2,1.4,6.8,2.7c1.7,1.2,3,7,1.2,7.7c-1.8,0.6-1.8,2.5,0.9,5.4 c2.7,2.9,3.4,6.8,4.2,7.5c0.4,0.3,0.9,0.5,0.9,0.5c0,0,7,1.1,8-1.8c0.9-2.9,3.4-4.1,6.1-5c2.7-0.8,4.1-1,2.3-4.4 c-1.9-3.3-1.4-3.5,0.8-5.8c2.3-2.3,4.5-3.7,6.4-3.7c1.9,0,6.6-1.2,7.3-3.3c0.6-2.1,0.2-2.5,4.3-2.9c4.1-0.4,6.2-4.6,6.6-6 c0.4-1.4,2.5-1.9,5-1.9c2.5,0,3.5-3.1,2.1-5c-1.5-1.9-1-3.9,3.7-3.3c4.8,0.6,4.6-3.1,8.3-2.7c1.5,0.2,3.3,0.5,5.1,0.8 C364.4,75.6,366.6,72.6,361.5,72.3z M313,75.4c-1.7,1,2.3,5.2,0,5.8c-7.7,0.4-8.1-3.1-11-5c-2.9-1.9-4.3-3.9-4.3-7 c0-3-1.4-5.5-5.6-8.4c-4.1-2.9-2.9-5.2-0.8-5.6c2.1-0.4,15-8.2,16.4-8.7c1.4-0.5,2.7,1.9,1.9,3.7c-0.8,1.9,0.2,3.9,1,7.7 c0.8,3.7,4.4,2.9,4.3,4.1c0,1.2,1.1,5.2,2.3,9.3C318.4,75.7,314.6,74.4,313,75.4z'
                }
            ]
		}
	},
    methods: {
        clickLocation(city){
            const target = city? {
                index: city.index,
                id: city.countycode,
            }: null
            this.$emit('trigger', target)
        } 
    }
}
</script>