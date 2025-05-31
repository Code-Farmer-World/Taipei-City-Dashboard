<template>
    <aside>
       
        <MapSelect 
            :location="location" 
            @trigger="clickLocation"
        />
        <div class="townBox">
            <TownSelect 
                v-if="location.index === 'taipei_city'"
                :town="tpTown" 
                @trigger="clickTpTown"
            />
            <TownSelectNewTaipei 
                v-if="location.index === 'new_taipei_city'" 
                :town="ntpTown" 
                @trigger="clickNTpTown"
            />
        </div>
        <MapChart 
            :location="location" 
            :tpTown="tpTown" 
            :ntpTown="ntpTown"
        />
    </aside>
    <MapBox 
        :location="location"
        :tp-town="tpTown"
        :ntp-town="ntpTown"
    /> 
</template>

<script>
import { defineAsyncComponent } from 'vue'
import MapSelect from "@/components/maps/MapSelect.vue"
import MapBox from "@/components/maps/MapBox.vue"
export default {
	data() {
		return {
            location: {},
            tpTown: {},
            ntpTown: {},
		}
	},
    computed: {
        currStep() {
			return this.$store.state.step
		},
        load() {
			return this.currStep >= 7 && this.currStep <= 9
		}
    },
	components:{
        MapSelect,
        MapBox,
		TownSelect: defineAsyncComponent(() => import('@/components/maps/TownSelect.vue')),
		TownSelectNewTaipei: defineAsyncComponent(() => import('@/components/maps/TownSelectNewTaipei.vue'))
	},
    methods: {
        clickLocation(target){
            this.location = target? target: {}
            this.tpTown = {}
            this.ntpTown = {}
        },
        clickTpTown(targetObj){
            this.tpTown = targetObj
            this.$gtag.event('update_map_filter', {
                event_category: 'neihu_traffic_action',
                value: 'tp'
            })
        },
        clickNTpTown(targetObj){
            this.ntpTown = targetObj
            this.$gtag.event('update_map_filter', {
                event_category: 'neihu_traffic_action',
                value: 'ntp'
            })
        } 
    }
}
</script>