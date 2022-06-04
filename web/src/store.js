import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const CLASSES = ['Mercenary', 'Medic', 'Hero', 'Engineer', 'Soldier', 'Ninja', 'Sniper', 'Scientist', 'Biologist', 'Looper'];

const store = new Vuex.Store({
  state: {
    playersDaily: [],
    dailyActive: [],
    weeklyActive: [],
    dailySurvivals: [],
    weeklySurvivals: [],
    classSurvivals: [],
    dailyKills: [],
    weeklyKills: [],
    online: [],
    activities: [],
    mvp: [],
    maps: [],
    classAliveTime: [],
    currentMap: 'infc_newdust',
    zombieWins: 0,
    humanWins: 0,
    picks: {},
    kills: {},
    ratio: {},
    records: {},
    score: {},
    best: {},
  },
  getters: {
    playerCount(state) {
      return state.online.length;
    },
    mapImage(state) {
      let img;
      try {
        let name = state.currentMap.split(" ")[0];
        img = require(`@/assets/maps/${name}.png`);
      } catch (ex) {
        return require(`@/assets/maps/infc_newdust.png`);
      }
      return img
    },
    // TODO: Pass getters
    classPicks: state => (className) => {
      const result = state.picks[className];
      return result === undefined ? [] : result;
    },
    classKills: state => (className) => {
      const result = state.kills[className];
      return result === undefined ? [] : result;
    },
    classRatio: state => (className) => {
      const result = state.ratio[className];
      return result === undefined ? [] : result;
    },
    classRecord: state => (className) => {
      const result = state.records[className];
      return result === undefined ? [] : result;
    },
    classBest: state => (className) => {
      const result = state.best[className];
      return result === undefined ? [] : result;
    },
    classTotalPicked: state => (className) => {
      const classPicks = state.picks[className];
      if (classPicks === undefined) {
        return 0;
      }

      return classPicks.reduce((sum, currentObj) => sum + currentObj.Value, 0);
    },
    classTotalScore: state => (className) => {
      const classScore = state.score[className];
      if (classScore === undefined) {
        return 0;
      }

      return classScore.reduce((sum, currentObj) => sum + currentObj.Value, 0);
    },
    classTotalKills: state => (className) => {
      const classKills = state.kills[className];
      if (classKills === undefined) {
        return 0;
      }

      return classKills.reduce((sum, currentObj) => sum + currentObj.Value, 0);
    },
    /**
     * Used for area-chart. Gives insight on how many players
     * played on the server at the time of the day
     * Apex-Charts wants specific objects see 2.1) of https://apexcharts.com/docs/series/
     */
    chartActivity(state) {
      const data = state.activities.map(activity => [activity.Timestamp, activity.Amount]);

      return [{
        name: 'Active players',
        // Convert aray of objects to array of arrays (which apexcharts understands)
        data,
      }];
    },
    chartAliveTime(state) {
      return [{
        name: 'series-survival',
        data: state.classAliveTime,
      }];
    },
    chartPlayerClass: state => (playerName) => {
      // Convert object with number properties to array of numbers
      // Puts the role-picked-amount for each class in an array
      const data = [];
      Object.keys(state.picks).forEach((role) => {
        const p = state.picks[role].find(obj => obj.Key === playerName);
        if (p === undefined) {
          data.push(0);
        } else {
          data.push(p.Value);
        }
      });
      return data;
    },
    chartPickDistribution(state, getters) {
      return CLASSES.map(role => getters.classTotalPicked(role));
    },
    chartKillDistribution(state, getters) {
      return CLASSES.map(role => getters.classTotalKills(role));
    },
    chartScoreDistribution(state, getters) {
      return CLASSES.map(role => getters.classTotalScore(role));
    },
    mapWin: state => (mapName, race) => {
      const m = state.maps.find(obj => obj.Name === mapName);
      return m === undefined ? 0 : m[race];
    },
    mapAverageSurvival: state => (mapName) => {
      const m = state.maps.find(obj => obj.Name === mapName);
      if (m === undefined) {
        return '0:00';
      }

      let seconds = Math.floor(m.Duration / m.Eligible);
      const minutes = Math.floor(seconds / 60);
      seconds -= minutes * 60;
      return `${minutes}:${(`00${seconds}`).substr(-2, 2)}`;
    },
    profileKills: state => (playerName) => {
      const p = state.weeklyKills.find(obj => obj.Key === playerName);
      return p === undefined ? '0' : p.Value;
    },
    profileTime: state => (playerName) => {
      const p = state.weeklyActive.find(obj => obj.Key === playerName);
      return p === undefined ? '0' : p.Value;
    },
  },
  mutations: {
    update(state, data) {
      state.currentMap = data.CurrentMap === 'UNKNOWN' ? 'infc_newdust' : data.CurrentMap;
      state.activities = data.Activities;
      state.dailyActive = data.DailyActive;
      state.dailyKills = data.DailyKills;
      state.dailySurvivals = data.DailySurvivals;
      state.weeklyActive = data.WeeklyActive;
      state.weeklyKills = data.WeeklyKills;
      state.weeklySurvivals = data.WeeklySurvivals;
      state.humanWins = data.HumanWins;
      state.zombieWins = data.ZombieWins;
      state.maps = data.Maps;
      state.mvp = data.Mvps;
      state.online = data.Online;
      state.picks = data.ClassPicks;
      state.kills = data.ClassKills;
      state.ratio = data.ClassRatio;
      state.records = data.ClassRecords;
      state.score = data.ClassScores;
      state.best = data.ClassBest;
      state.classSurvivals = data.ClassSurvivals;
      state.classAliveTime = data.ClassAliveTime;
    },
  },
});

export default store;