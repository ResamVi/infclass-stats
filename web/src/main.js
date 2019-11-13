/*
    Website start
    Copyright (C) 2018  Julien Midedji

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import Vue from 'vue';
import Router from 'vue-router';
import { Tabs, Tab } from 'vue-tabs-component';
import VueApexCharts from 'vue-apexcharts';

import App from './App.vue';
import Home from './views/Home.vue';
import About from './views/About.vue';
import Profile from './views/Profile.vue';
import Search from './views/Search.vue';
import store from './store';

Vue.config.productionTip = false;

Vue.use(VueApexCharts);
Vue.use(Router);

Vue.component('apexchart', VueApexCharts);
Vue.component('tabs', Tabs);
Vue.component('tab', Tab);

const router = new Router({
  routes: [
    { path: '/', name: 'home', component: Home },
    { path: '/about', name: 'about', component: About },
    { path: '/search', name: 'search', component: Search },
    { path: '/player/:name', name: 'player', component: Profile },
  ],
});

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
