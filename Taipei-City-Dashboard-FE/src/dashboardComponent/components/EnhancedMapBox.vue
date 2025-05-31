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
const mapConfig = {
  container: 'enhanced-mapbox-container',
  style: 'mapbox://styles/mapbox/light-v10',
  center: [121.5654, 25.0330], // Taipei center
  zoom: 10,
  pitch: 45,
  bearing: 0,
  antialias: true
}

// Enhanced style configuration based on neihu_traffic
const enhancedStyle = {
  PopWorkStyle: {
    'fill-extrusion-color': [
      'interpolate',
      ['linear'],
      [
        '+',
        ['to-number', ['get', '0歲數量'] || 0],
        ['to-number', ['get', '1歲數量'] || 0],
        ['to-number', ['get', '2歲數量'] || 0],
        ['to-number', ['get', '3歲數量'] || 0],
        ['to-number', ['get', '4歲數量'] || 0],
        ['to-number', ['get', '5歲數量'] || 0]
      ],
      0, '#0C303E',
      50, '#184556',
      100, '#1E4F63',
      150, '#24596F',
      200, '#2A647B',
      250, '#2F6E87',
      300, '#3B829F',
      350, '#4797B7',
      400, '#52ABCF',
      450, '#5EC0E7'
    ],
    'fill-extrusion-height': [
      'interpolate',
      ['linear'],
      [
        '+',
        ['to-number', ['get', '0歲數量'] || 0],
        ['to-number', ['get', '1歲數量'] || 0],
        ['to-number', ['get', '2歲數量'] || 0],
        ['to-number', ['get', '3歲數量'] || 0],
        ['to-number', ['get', '4歲數量'] || 0],
        ['to-number', ['get', '5歲數量'] || 0]
      ],
      0, 0,
      50, 100,
      100, 200,
      150, 300,
      200, 400,
      250, 500,
      300, 600,
      350, 700,
      400, 800,
      450, 900,
      500, 1000
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

// Load map data
const loadMapData = async () => {
  try {
    // Load small.geojson data
    const response = await fetch('https://github.com/Code-Farmer-World/Taipei-City-Dashboard/blob/01aab7d995948e591e20ec2dce08515044d5102c/neihu_traffic/public/data/small.geojson')
    const geojsonData = await response.json()
    
    // Add data source
    mapInstance.value.addSource('newborn-data', {
      type: 'geojson',
      data: geojsonData
    })
    
    // Add 3D buildings layer
    const layers = mapInstance.value.getStyle().layers
    const labelLayerId = layers.find(
      (layer) => layer.type === 'symbol' && layer.layout['text-field']
    ).id
    
    // Add enhanced 3D extrusion layer
    mapInstance.value.addLayer({
      id: 'newborn-3d',
      type: 'fill-extrusion',
      source: 'newborn-data',
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
  mapInstance.value.on('click', 'newborn-3d', (e) => {
    const coordinates = e.lngLat
    const properties = e.features[0].properties
    
    showPopup(coordinates, properties)
  })
  
  // Hover effects
  mapInstance.value.on('mouseenter', 'newborn-3d', () => {
    mapInstance.value.getCanvas().style.cursor = 'pointer'
  })
  
  mapInstance.value.on('mouseleave', 'newborn-3d', () => {
    mapInstance.value.getCanvas().style.cursor = ''
  })
}

// Show popup with data
const showPopup = (coordinates, properties) => {
  if (popup.value) {
    popup.value.remove()
  }
  
  const totalPopulation = (
    parseInt(properties['0歲數量'] || 0) +
    parseInt(properties['1歲數量'] || 0) +
    parseInt(properties['2歲數量'] || 0) +
    parseInt(properties['3歲數量'] || 0) +
    parseInt(properties['4歲數量'] || 0) +
    parseInt(properties['5歲數量'] || 0)
  )
  
  const popupContent = `
    <div class="mapbox-popup">
      <h3>${properties.VILLNAME || '未知地區'}</h3>
      <p><strong>行政區：</strong>${properties.TOWNNAME || '未知'}</p>
      <p><strong>總幼兒人數：</strong>${totalPopulation} 人</p>
      <div class="age-breakdown">
        <p>0歲：${properties['0歲數量'] || 0} 人</p>
        <p>1歲：${properties['1歲數量'] || 0} 人</p>
        <p>2歲：${properties['2歲數量'] || 0} 人</p>
        <p>3歲：${properties['3歲數量'] || 0} 人</p>
        <p>4歲：${properties['4歲數量'] || 0} 人</p>
        <p>5歲：${properties['5歲數量'] || 0} 人</p>
      </div>
    </div>
  `
  
  popup.value = new mapboxgl.Popup()
    .setLngLat(coordinates)
    .setHTML(popupContent)
    .addTo(mapInstance.value)
}

// Watch for data changes
watch(() => props.series, () => {
  if (mapInstance.value && mapInstance.value.isStyleLoaded()) {
    // Update map data when series changes
    loadMapData()
  }
}, { deep: true })

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
  height: 400px;
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

.age-breakdown {
  margin-top: 10px;
  padding-top: 10px;
  border-top: 1px solid #eee;
}

.age-breakdown p {
  margin: 3px 0;
  font-size: 12px;
}

.mapboxgl-popup-content {
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.mapboxgl-popup-tip {
  border-top-color: white;
}
</style>