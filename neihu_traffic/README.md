#### geojson 換成讀取 api

`loadDataToMapbox` 這方法要讀取api回傳geojson格式
分別queryString 帶入taipei,both_taipei

geojson 範例使用small.geojson
否則資料量過大

```
{
"type": "FeatureCollection",
"name": "neihu_pop",
"crs": { "type": "name", "properties": { "name": "urn:ogc:def:crs:OGC:1.3:CRS84" } },
"features": [
{ "type": "Feature", "properties": { "id": "52649", "pop_work_min": 0, "bus_up": 0, "bus_down": 0, "mrt_up": 0, "mrt_down": 0, "ubike_up": 0, "ubike_down": 0, "VILLCODE": "63000120038", "COUNTYNAME": "臺北市", "TOWNNAME": "北投區", "VILLNAME": "關渡里", "VILLENG": "Guandu Vil.", "COUNTYID": "A", "COUNTYCODE": "63000", "TOWNID": "A16", "TOWNCODE": "63000120", "bus_avg": 0, "mrt_avg": 0, "ubike_avg": 0, "transport_avg": 0, "untransport": 0, "transport_rate": null }, "geometry": { "type": "Polygon", "coordinates": [ [ [ 121.460335350924879, 25.122093014058073 ], [ 121.460343807758505, 25.124349999106137 ], [ 121.462822946893809, 25.124342278891071 ], [ 121.462814444520646, 25.122085294630967 ], [ 121.460335350924879, 25.122093014058073 ] ] ] } }
]
}
```

#### 物件說明 可以在圖磚上顯示訊息 顯示業務欄位

```

langObj : {
                COUNTYNAME: '來源地縣市',
                TOWNNAME: '行政區',
                // VILLNAME: '里',
                //COUNTYCODE - location.id
                //TOWNID - tpTown.id || ntpTown.id
                pop_work_min: '總工作人口',
                transport_avg: '大眾運輸通勤人口',
                transport_rate: '使用大眾運輸比例',
                // bus_avg: '公車轉乘',
                // mrt_avg: '捷運轉乘',
                // ubike_avg: '腳踏車轉乘',
                // untransport: '私人運具',
            }
```
透過coordinates渲染形成圖磚或其他形狀
"type": "Polygon"



