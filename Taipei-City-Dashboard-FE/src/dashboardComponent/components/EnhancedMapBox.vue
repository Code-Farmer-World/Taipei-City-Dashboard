<!-- Enhanced MapBox Component with 3D Effects -->
<!-- Based on neihu_traffic MapBox implementation -->

<template>
  <div class="enhanced-mapbox">
    <div id="enhanced-mapbox-container" class="mapbox-container" />
    <div v-if="loading" class="loading-overlay">
      <div class="loading-spinner" />
      <p>載入地圖中...</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch, nextTick } from 'vue'
import mapboxgl from 'mapbox-gl'
import 'mapbox-gl/dist/mapbox-gl.css'

// Props
const props = defineProps({
  series: {
    type: Array,
    default: () => []
  },
  map_config: {
    type: Array,
    default: () => []
  },
  map_filter: {
    type: Object,
    default: () => ({})
  },
  chart_config: {
    type: Object,
    default: () => ({})
  },
  activeChart: {
    type: String,
    default: 'MapLegend'
  },
  activeCity: {
    type: String,
    default: ''
  },
  // 新增地圖區域切換 prop
  mapRegion: {
    type: String,
    default: 'taipei', // 'taipei', 'metrotaipei'
    validator: (value) => ['taipei', 'metrotaipei'].includes(value)
  }
})

// Emits
const emit = defineEmits([
  'filterByParam',
  'filterByLayer',
  'clearByParamFilter',
  'clearByLayerFilter',
  'fly'
])

// Reactive data
const loading = ref(true)
const mapInstance = ref(null)
const popup = ref(null)
const currentFilter = ref(null)

// Mapbox configuration
const MAPBOXTOKEN = 'pk.eyJ1Ijoia2trMTIzNTUiLCJhIjoiY21hdXN0ZzQxMDBocjJtcHA0bGFla2xjYyJ9.VT_ubDB8ck90VbCCz4HdHg'

// 地圖區域配置
const mapRegionConfig = {
  taipei: {
    center: [121.55585298158064, 25.05244617333119],
    zoom: 11.5,
    pitch: 45,
    bearing: 0,
    bounds: [
      [121.40842683023203, 24.958170494168186],
      [121.69688923958796, 25.21839405253573]
    ]
  },
  metrotaipei: {
    center: [121.5394269250055, 25.037188097677472],
    zoom: 10.5,
    pitch: 45,
    bearing: 0,
    bounds: [
      [121.0471573380878, 24.695612549499685],
      [121.92655852488502, 25.30785665565]
    ]
  }
}

// 動畫配置
const animationConfig = {
  duration: 3000, // 3秒動畫
  easing: 'ease-in-out'
}

const mapConfig = {
  container: 'enhanced-mapbox-container',
  style: 'mapbox://styles/mapbox/light-v10',
  center: mapRegionConfig.taipei.center,
  zoom: mapRegionConfig.taipei.zoom,
  pitch: mapRegionConfig.taipei.pitch,
  bearing: mapRegionConfig.taipei.bearing,
  antialias: true
}

// Enhanced style configuration based on neihu_traffic
const enhancedStyle = {
  PopWorkStyle: {
    'fill-extrusion-color': [
      'interpolate',
      ['linear'],
      ['to-number', ['get', 'pop_work_min'], 0],
      0, '#0C303E',
      10, '#184556',
      20, '#1E4F63',
      30, '#24596F',
      40, '#2A647B',
      50, '#2F6E87',
      60, '#3B829F',
      70, '#4797B7',
      80, '#52ABCF',
      90, '#5EC0E7'
    ],
    'fill-extrusion-height': [
      'interpolate',
      ['linear'],
      ['to-number', ['get', 'pop_work_min'], 0],
      0, 0,
      10, 200,
      20, 400,
      30, 600,
      40, 800,
      50, 1000,
      60, 1200,
      70, 1400,
      80, 1600,
      90, 1800,
      100, 2000
    ],
    'fill-extrusion-opacity': 0.8
  },
  mapboxBuildings: {
    'fill-extrusion-color': '#aaa',
    'fill-extrusion-height': [
      'interpolate',
      ['linear'],
      ['zoom'],
      15, 0,
      15.05, ['get', 'height']
    ],
    'fill-extrusion-base': [
      'interpolate',
      ['linear'],
      ['zoom'],
      15, 0,
      15.05, ['get', 'min_height']
    ],
    'fill-extrusion-opacity': 0.6
  }
}

// Initialize Mapbox
const initializeMap = async () => {
  try {
    loading.value = true
    
    // Use the provided Mapbox token
    mapboxgl.accessToken = MAPBOXTOKEN
    
    mapInstance.value = new mapboxgl.Map(mapConfig)
    
    mapInstance.value.on('load', async () => {
      await loadMapData()
      setupMapInteractions()
      loading.value = false
    })
    
    mapInstance.value.on('error', (e) => {
      console.error('Mapbox error:', e)
      console.error('Error details:', e.error)
      // Try to use a different style if the current one fails
      if (e.error && e.error.status === 403) {
        console.warn('Access denied, trying alternative map style...')
        tryAlternativeMapStyle()
      } else {
        loading.value = false
      }
    })
    
    // Add navigation controls
    mapInstance.value.addControl(new mapboxgl.NavigationControl())
    
  } catch (error) {
    console.error('Failed to initialize map:', error)
    loading.value = false
  }
}

// Try alternative map style if main style fails
const tryAlternativeMapStyle = () => {
  try {
    // Use a simpler, more accessible style
    mapInstance.value.setStyle('mapbox://styles/mapbox/streets-v11')
    mapInstance.value.once('styledata', async () => {
      await loadMapData()
      setupMapInteractions()
      loading.value = false
    })
  } catch (error) {
    console.error('Alternative style also failed:', error)
    loading.value = false
  }
}

// Load map data based on region
const loadMapData = async () => {
  try {
    // Determine which data file to load based on mapRegion
    let dataFile = '/data/small.geojson' // default for taipei
    
    if (props.mapRegion === 'metrotaipei') {
      // For Metro Taipei (both cities), use the combined data
      dataFile = '/mapData/metrotaipei_village.geojson'
    } else {
      // For Taipei City, use the existing small.geojson
      dataFile = '/data/small.geojson'
    }
    
    const response = await fetch(dataFile)
    const geojsonData = await response.json()
    
    // Filter data based on region if needed
    let filteredData = geojsonData
    if (props.mapRegion === 'taipei' && geojsonData.features && dataFile.includes('metrotaipei')) {
      // Filter for Taipei City only when using metrotaipei data
      filteredData = {
        ...geojsonData,
        features: geojsonData.features.filter(feature => 
          feature.properties && 
          (feature.properties.PNAME === '臺北市' || feature.properties.COUNTYNAME === '臺北市')
        )
      }
    }
    // For metrotaipei, use all data without filtering
    
    // Remove existing source if it exists
    if (mapInstance.value.getSource('senior-service-data')) {
      mapInstance.value.removeLayer('senior-service-3d')
      mapInstance.value.removeSource('senior-service-data')
    }
    
    // Add data source
    mapInstance.value.addSource('senior-service-data', {
      type: 'geojson',
      data: filteredData
    })
    
    // Add 3D buildings layer
    const layers = mapInstance.value.getStyle().layers
    const labelLayerId = layers.find(
      (layer) => layer.type === 'symbol' && layer.layout['text-field']
    ).id
    
    // Add enhanced 3D extrusion layer
    mapInstance.value.addLayer({
      id: 'senior-service-3d',
      type: 'fill-extrusion',
      source: 'senior-service-data',
      layout: {},
      paint: enhancedStyle.PopWorkStyle
    }, labelLayerId)
    
    // Add building layer for context
    mapInstance.value.addLayer({
      id: 'building-3d',
      source: 'composite',
      'source-layer': 'building',
      filter: ['==', 'extrude', 'true'],
      type: 'fill-extrusion',
      minzoom: 15,
      paint: enhancedStyle.mapboxBuildings
    })
    
  } catch (error) {
    console.error('Failed to load map data:', error)
  }
}

// Setup map interactions
const setupMapInteractions = () => {
  // Click event for popup
  mapInstance.value.on('click', 'senior-service-3d', (e) => {
    const coordinates = e.lngLat
    const properties = e.features[0].properties
    
    showPopup(coordinates, properties)
  })
  
  // Hover effects
  mapInstance.value.on('mouseenter', 'senior-service-3d', () => {
    mapInstance.value.getCanvas().style.cursor = 'pointer'
  })
  
  mapInstance.value.on('mouseleave', 'senior-service-3d', () => {
    mapInstance.value.getCanvas().style.cursor = ''
  })
}

// Show popup with data
const showPopup = (coordinates, properties) => {
  if (popup.value) {
    popup.value.remove()
  }
  
  const workingPopulation = parseInt(properties.pop_work_min || 0)
  const transportAvg = parseInt(properties.transport_avg || 0)
  const transportRate = properties.transport_rate ? (properties.transport_rate * 100).toFixed(1) : '無資料'
  
  const popupContent = `
    <div class="mapbox-popup">
      <h3>${properties.VILLNAME || '未知地區'}</h3>
      <p><strong>行政區：</strong>${properties.TOWNNAME || '未知'}</p>
      <p><strong>工作人口數：</strong>${workingPopulation} 人</p>
      <p><strong>交通便利度：</strong>${transportRate}%</p>
      <div class="service-breakdown">
        <h4>銀髮族交通服務便利性分析</h4>
        <p>🚌 公車服務：上車 ${properties.bus_up || 0} 人次 / 下車 ${properties.bus_down || 0} 人次</p>
        <p>🚇 捷運服務：上車 ${properties.mrt_up || 0} 人次 / 下車 ${properties.mrt_down || 0} 人次</p>
        <p>🚲 Ubike服務：借車 ${properties.ubike_up || 0} 人次 / 還車 ${properties.ubike_down || 0} 人次</p>
        <p>📊 平均交通使用量：${transportAvg} 人次</p>
        <p>🏠 非交通通勤人口：${properties.untransport || 0} 人</p>
        <p class="service-note">💡 此區域銀髮族可透過多元交通工具便利出行</p>
      </div>
    </div>
  `
  
  popup.value = new mapboxgl.Popup()
    .setLngLat(coordinates)
    .setHTML(popupContent)
    .addTo(mapInstance.value)
}

// 地圖區域切換方法
const switchMapRegion = async (region) => {
  if (!mapInstance.value || !mapRegionConfig[region]) return
  
  const config = mapRegionConfig[region]
  
  // Reload map data for the new region
  await loadMapData()
  
  // Re-add the 3D layer after loading new data
  const layers = mapInstance.value.getStyle().layers
  const labelLayerId = layers.find(
    (layer) => layer.type === 'symbol' && layer.layout['text-field']
  )?.id
  
  if (labelLayerId && !mapInstance.value.getLayer('senior-service-3d')) {
    mapInstance.value.addLayer({
      id: 'senior-service-3d',
      type: 'fill-extrusion',
      source: 'senior-service-data',
      layout: {},
      paint: enhancedStyle.PopWorkStyle
    }, labelLayerId)
  }
  
  // 使用 easeTo 進行平滑動畫切換
  mapInstance.value.easeTo({
    center: config.center,
    zoom: config.zoom,
    pitch: config.pitch,
    bearing: config.bearing,
    duration: animationConfig.duration,
    easing: animationConfig.easing
  })
  
  // 可選：同時調整地圖邊界
  setTimeout(() => {
    if (config.bounds) {
      mapInstance.value.fitBounds(config.bounds, {
        duration: animationConfig.duration / 2,
        padding: 20
      })
    }
  }, animationConfig.duration / 2)
}

// Watch for data changes
watch(() => props.series, () => {
  if (mapInstance.value && mapInstance.value.isStyleLoaded()) {
    // Update map data when series changes
    loadMapData()
  }
}, { deep: true })

// Watch for map region changes
watch(() => props.mapRegion, (newRegion) => {
  if (mapInstance.value && mapInstance.value.isStyleLoaded()) {
    switchMapRegion(newRegion)
  }
}, { immediate: false })

// Watch for activeCity changes and map to region
watch(() => props.activeCity, async (newCity) => {
  if (mapInstance.value && mapInstance.value.isStyleLoaded()) {
    // Map activeCity values to mapRegion values
    let targetRegion = 'taipei' // default
    
    if (newCity === 'taipei') {
      targetRegion = 'taipei'
    } else if (newCity === 'metrotaipei') {
      targetRegion = 'metrotaipei'
    }
    
    // Update the mapRegion prop to trigger data reload
    const currentMapRegion = props.mapRegion
    if (currentMapRegion !== targetRegion) {
      // Temporarily update mapRegion for data loading
      Object.defineProperty(props, 'mapRegion', {
        value: targetRegion,
        writable: true,
        configurable: true
      })
    }
    
    await switchMapRegion(targetRegion)
  }
}, { immediate: false })

// 暴露方法給父組件使用
defineExpose({
  switchMapRegion,
  mapInstance: () => mapInstance.value
})

// Lifecycle hooks
onMounted(() => {
  nextTick(() => {
    initializeMap()
  })
})

onUnmounted(() => {
  if (popup.value) {
    popup.value.remove()
  }
  if (mapInstance.value) {
    mapInstance.value.remove()
  }
})
</script>

<style scoped>
.enhanced-mapbox {
  position: relative;
  width: 100%;
  height: 100%;
  border-radius: 5px;
  overflow: hidden;
}

.mapbox-container {
  width: 100%;
  height: 100%;
}

.loading-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: white;
  z-index: 1000;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-top: 3px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 10px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>

<style>
/* Global popup styles */
.mapbox-popup {
  font-family: inherit;
}

.mapbox-popup h3 {
  margin: 0 0 10px 0;
  color: #333;
  font-size: 16px;
}

.mapbox-popup p {
  margin: 5px 0;
  font-size: 14px;
  color: #666;
}

.service-breakdown {
  margin-top: 10px;
  padding-top: 10px;
  border-top: 1px solid #eee;
}

.service-breakdown p {
  margin: 3px 0;
  font-size: 12px;
}

.service-note {
  margin-top: 8px;
  padding: 5px;
  background-color: #f0f8ff;
  border-radius: 4px;
  font-style: italic;
  color: #2c5aa0;
}

.mapboxgl-popup-content {
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.mapboxgl-popup-tip {
  border-top-color: white;
}
</style>