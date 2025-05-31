import { createApp } from 'vue'
import { createStore } from 'vuex'

import App from './App.vue'

import { createI18n } from 'vue-i18n'
import { en, zh } from './assets/js/lang.js'

import VueApexCharts from "vue3-apexcharts"

import VueGtag from "vue-gtag"

import AnnotationTrigger from "@/components/annotation/Trigger.vue"
import AnnotationContent from "@/components/annotation/Content.vue"

import '@/assets/scss/reset.scss'

const Locale = localStorage.getItem("locale")
const i18n = createI18n({
    locale: Locale ? Locale: "zh-TW",
    messages: {
		"zh-TW": zh,
		"en-US": en
    }
})

const store = createStore({
	state () {
		return {
			step: 0,
			progres: 0,
			contentEnter: false,
			mobileDevice: false,
			videoStart: false,
			lowSpeed: false
		}
	},
	mutations: {
		updateStep (state, payload) {
			state.step = payload
		},
		updateProgres (state, payload) {
			state.progres = payload
		},
		updateContentEnter (state, payload) {
			state.contentEnter = payload
		},
		updateMobileDevice (state, payload) {
			state.mobileDevice = payload
		},
		toggleVideoStatus (state, payload) {
			state.videoStart = payload
		},
		updateLowSpeed (state, payload) {
			state.lowSpeed = payload
		}
	}
})

const app = createApp(App)
app.use(i18n)
app.use(VueApexCharts)
app.use(store)
app.use(VueGtag, {
	appName: 'neihu_traffic',
	pageTrackerScreenviewEnabled: true,
	config: { id: process.env.VUE_APP_GA_MEASUREMENT_ID }
})
app.component('AnnotationTrigger', AnnotationTrigger)
app.component('AnnotationContent', AnnotationContent)
app.mount('#app')
