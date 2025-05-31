<!-- Enhanced MapBox Component with 3D Effects -->
<!-- Based on neihu_traffic MapBox implementation -->

<template>
  <div class="enhanced-mapbox" :class="{ 'dark': isDarkTheme }">
    <div id="enhanced-mapbox-container" class="mapbox-container" ref="mapContainer" />
    <div v-if="loading" class="loading-overlay">
      <div class="loading-spinner" />
      <p>載入地圖中...</p>
    </div>
    <!-- 主題切換按鈕 -->
    <div class="theme-toggle-container">
      <button
        class="theme-toggle-btn"
        @click="toggleMapTheme"
        :title="isDarkTheme ? '切換到亮色主題' : '切換到暗色主題'"
      >
        <span class="material-icons">{{ isDarkTheme ? 'light_mode' : 'dark_mode' }}</span>
      </button>
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
const isDarkTheme = ref(false)

// Mapbox configuration
const MAPBOXTOKEN = 'pk.eyJ1Ijoia2trMTIzNTUiLCJhIjoiY21hdXN0ZzQxMDBocjJtcHA0bGFla2xjYyJ9.VT_ubDB8ck90VbCCz4HdHg'

// 地圖主題配置
const mapThemes = {
  light: 'mapbox://styles/mapbox/light-v10',
  dark: 'mapbox://styles/mapbox/dark-v10'
}

// 地圖區域配置
const mapRegionConfig = {
  taipei: {
    center: [121.46827859298899, 25.120812434425802], // 調整到關渡地區
    zoom: 16, // 增加縮放級別以便看到小區域
    pitch: 45,
    bearing: 0,
    bounds: [
      [121.40842683023203, 24.958170494168186],
      [121.69688923958796, 25.21839405253573]
    ]
  },
  metrotaipei: {
    center: [121.46827859298899, 25.120812434425802], // 同樣調整到關渡地區
    zoom: 15,
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

const getMapConfig = () => ({
  container: 'enhanced-mapbox-container',
  style: isDarkTheme.value ? mapThemes.dark : mapThemes.light,
  center: mapRegionConfig.taipei.center,
  zoom: mapRegionConfig.taipei.zoom,
  pitch: mapRegionConfig.taipei.pitch,
  bearing: mapRegionConfig.taipei.bearing,
  antialias: true
})

// Enhanced style configuration for senior service data
const enhancedStyle = {
  PopWorkStyle: {
    'fill-extrusion-color': [
      'interpolate',
      ['linear'],
      ['to-number', ['get', 'annual_expected_participants'], 0],
      0, '#0C303E',
      20, '#184556',
      40, '#1E4F63',
      60, '#24596F',
      80, '#2A647B',
      100, '#2F6E87',
      120, '#3B829F',
      140, '#4797B7',
      160, '#52ABCF',
      180, '#5EC0E7'
    ],
    'fill-extrusion-height': [
      'interpolate',
      ['linear'],
      ['to-number', ['get', 'annual_expected_participants'], 0],
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

    mapInstance.value = new mapboxgl.Map(getMapConfig())

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
    // Get data from API
    const city = props.activeCity || props.mapRegion || 'taipei'
	const apiUrl = `/data/small2.geojson`

	
    
    const response = await fetch(apiUrl)
     const apiData = await response.json()
    
    // Check if API response has valid data
    if (!apiData || !apiData.features || !Array.isArray(apiData.features)) {
        console.warn('Invalid API response structure:', apiData)
        return
      }
    
    // Convert API data to GeoJSON format
    // API 回應已經是標準的 GeoJSON 格式，直接使用
    const geojsonData = {
      type: 'FeatureCollection',
      features: apiData.features.map(feature => ({
        ...feature,
        properties: {
          ...feature.properties,
          // Add random properties for 3D effect
          height: Math.random() * 200 + 50,
          base_height: Math.random() * 50
        }
      }))
    }

    // Use all data without filtering for now to ensure visibility
    let filteredData = geojsonData
    
    console.log('Original geojsonData:', geojsonData)
    console.log('Number of features:', geojsonData.features.length)
    console.log('First feature properties:', geojsonData.features[0]?.properties)
    console.log('First feature coordinates:', geojsonData.features[0]?.geometry?.coordinates)
    
    // Temporarily disable filtering to debug rendering issues
    // if (props.mapRegion === 'taipei') {
    //   // Filter for Taipei City only - based on coordinates
    //   filteredData = {
    //     ...geojsonData,
    //     features: geojsonData.features.filter(feature =>
    //       feature.geometry && feature.geometry.coordinates && 
    //       feature.geometry.coordinates[0] && feature.geometry.coordinates[0][0] &&
    //       feature.geometry.coordinates[0][0][0] >= 121.4 && feature.geometry.coordinates[0][0][0] <= 121.7 &&
    //       feature.geometry.coordinates[0][0][1] >= 25.0 && feature.geometry.coordinates[0][0][1] <= 25.3
    //     )
    //   }
    // }

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

  const expectedParticipants = parseInt(properties.annual_expected_participants || 0)
  const districtRatio = properties.annual_district_ratio ? (properties.annual_district_ratio * 100).toFixed(1) : '無資料'
  const courseStatus = properties.course_status || '未知'
  const isPopular = properties.popular_course === 'true' || properties.popular_course === true
  const provideMeal = properties.provide_meal === 'true' || properties.provide_meal === true

  const popupContent = `
    <div class="mapbox-popup">
      <h3>${properties.course || '未知課程'}</h3>
      <p><strong>機構名稱：</strong>${properties.org_name || '未知'}</p>
      <p><strong>城市：</strong>${properties.city || '未知'}</p>
      <p><strong>行政區：</strong>${properties.district || '未知'}</p>
      <p><strong>地址：</strong>${properties.address || '未知'}</p>
      <div class="service-breakdown">
        <h4>課程詳細資訊</h4>
        <p>📚 課程類別：${properties.category || '未分類'}</p>
        <p>👥 預期參與人數：${expectedParticipants} 人</p>
        <p>📊 區域比例：${districtRatio}%</p>
        <p>📋 課程狀態：${courseStatus}</p>
        <p>🍽️ 提供餐點：${provideMeal ? '是' : '否'}</p>
        <p>⭐ 熱門課程：${isPopular ? '是' : '否'}</p>
        <p class="service-note">💡 ${isPopular ? '這是一個熱門的銀髮族課程！' : '歡迎參與銀髮族課程活動'}</p>
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

// 主題切換函數
const toggleMapTheme = () => {
  isDarkTheme.value = !isDarkTheme.value

  if (mapInstance.value && mapInstance.value.isStyleLoaded()) {
    const newStyle = isDarkTheme.value ? mapThemes.dark : mapThemes.light

    // 保存當前的地圖狀態
    const currentCenter = mapInstance.value.getCenter()
    const currentZoom = mapInstance.value.getZoom()
    const currentPitch = mapInstance.value.getPitch()
    const currentBearing = mapInstance.value.getBearing()

    // 切換地圖樣式
    mapInstance.value.setStyle(newStyle)

    // 等待樣式載入完成後重新添加數據層
    mapInstance.value.once('styledata', async () => {
      // 恢復地圖狀態
      mapInstance.value.setCenter(currentCenter)
      mapInstance.value.setZoom(currentZoom)
      mapInstance.value.setPitch(currentPitch)
      mapInstance.value.setBearing(currentBearing)

      // 重新載入地圖數據和圖層
      await loadMapData()
      setupMapInteractions()
    })
  }
}

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

@import url('https://fonts.googleapis.com/icon?family=Material+Icons');

/* ────────── 地圖容器 ────────── */
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

/* ────────── 載入遮罩 ────────── */
.loading-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #fff;
  z-index: 1000;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-top: 3px solid #fff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 10px;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* ────────── 主題切換按鈕 ────────── */
.theme-toggle-container {
  position: absolute;
  top: 110px;
  right: 5px;
  z-index: 1000;
}

/* ✔ 精簡後唯一的 `.theme-toggle-btn` 定義 */
.theme-toggle-btn {
  width: 40px;
  height: 40px;
  padding: 0;
  border: 1px solid #e0e0e0;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.9);
  color: #333;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  cursor: pointer;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.theme-toggle-btn:hover {
  background: #f0f0f0;

  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.theme-toggle-btn:active {
  transform: translateY(0);
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
}

.theme-toggle-btn span,
.theme-toggle-btn .icon {
  font-size: 20px;
  color: #333; /* 預設深色 */
  transition: all 0.3s ease;
}



/* ────────── 暗色主題覆寫 ────────── */
.enhanced-mapbox.dark .theme-toggle-btn {
  background: rgba(45, 45, 45, 0.9);
  border-color: #555;
}

.enhanced-mapbox.dark .theme-toggle-btn .icon,
.enhanced-mapbox.dark .theme-toggle-btn span {
  color: #fff !important;
}

.enhanced-mapbox.dark .theme-toggle-btn:hover {
  background: rgba(60, 60, 60, 0.95);
}

/* ────────── Popup 文字樣式 ────────── */
.mapbox-popup {
  font-family: inherit;
}

.mapbox-popup h3 {
  margin: 0 0 10px;
  color: #333;
  font-size: 16px;
}

.mapbox-popup p {
  margin: 5px 0;
  font-size: 14px;
  color: #666;
}

/* 服務細項 */
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
  background: #f0f8ff;
  border-radius: 4px;
  font-style: italic;
  color: #2c5aa0;
}

/* Mapbox 原生 Popup 美化 */
.mapboxgl-popup-content {
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.mapboxgl-popup-tip {
  border-top-color: #fff;
}

.material-icons {
  font-family: 'Material Icons';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;
  line-height: 1;
  letter-spacing: normal;
  text-transform: none;
  display: inline-block;
  white-space: nowrap;
  word-wrap: normal;
  direction: ltr;
  -webkit-font-feature-settings: 'liga';
  -webkit-font-smoothing: antialiased;
  color: inherit;
}
</style>

