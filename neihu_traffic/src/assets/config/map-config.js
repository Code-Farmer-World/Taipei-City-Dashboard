export const locationsCenter = {
    taiwan: [120.733393686039, 24.196575671135506],
    haveData: [121.5394269250055, 25.037188097677472],
    northArea: [121.4936211022997, 24.977560206267384],
    taipei: [121.55585298158064, 25.05244617333119],
    new_taipei: [121.555998642252, 24.946784805972158],
    keelung: [121.71505, 25.11795],
    taoyuan: [121.181306, 24.974826]
}

export const initZoom = {
    taiwan:7,
    haveData: 10.5,
    northArea: 9.75,
    taipeiArea: 10.75,
    taipei: 11.5,
    new_taipei:10,
    keelung: 12,
    taoyuan: 11,
    heatmp: 16.5,
    maxZoom: 18
}

/** maxBound
 *  [Southwest coordinates, Northeast coordinates]
 */
export const maxBound = {
    taiwan: [
        [118.06372184152644, 21.810549724023645],
        [122.62304381162107, 26.371897242606593]
    ],
    northArea: [
        [121.0471573380878, 24.695612549499685],
        [121.92655852488502, 25.30785665565]
    ],
    new_taipei: [
        [121.26822598517884, 24.647723381441295],
        [122.0187784250956, 25.325357955401643]
    ],
    taipei: [
        [121.40842683023203, 24.958170494168186],
        [121.69688923958796, 25.21839405253573]
    ],
    keelung: [
        [121.61932141997676, 25.039389264156412],
        [121.80901839903697, 25.21046481675478]
    ],
    taoyuan: [
        [120.97808316593915, 24.769783329104555],
        [121.4216970676469, 25.170354958418102]
    ]
}

export const durationConfig = 5000
export const fitPadding = {
    top: 5, 
    bottom:5, 
    left: 5, 
    right: 5
}
export const fitBoundsConfig = {
    padding: fitPadding,
    duration: durationConfig,
    maxZoom: initZoom.maxZoom + 1
}
