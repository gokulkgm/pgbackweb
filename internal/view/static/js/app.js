import { initThemeHelper } from './init-theme-helper.js'
import { initSweetAlert2 } from './init-sweetalert2.js'
import { initNotyf } from './init-notyf.js'
import { initHTMX } from './init-htmx.js'
import { initAlpineComponents } from './init-alpine-components.js'
import { initHelpers } from './init-helpers.js'
import { initDashboardAsideScroll } from './dashboard-aside-scroll.js'

initThemeHelper()
initSweetAlert2()
initNotyf()
initHTMX()
initAlpineComponents()
initHelpers()
initDashboardAsideScroll()

// Add event listeners to handle mobile responsiveness
window.addEventListener('resize', handleResize)
window.addEventListener('orientationchange', handleResize)

function handleResize() {
  const isMobile = window.innerWidth <= 600
  document.body.classList.toggle('mobile', isMobile)
}

// Initial check
handleResize()
