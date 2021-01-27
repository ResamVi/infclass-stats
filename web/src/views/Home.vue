<template>
  <main>
    <!-- MVP -->
    <article id="mvp">
      <Table
        :title="'BEST PLAYERS'"
        :area="'mvp'"
        :column="'Rating'"
        :list="mvp"
        :format="identity"
      />
    </article>

    <!-- Activity stats -->
    <article id="activity">
      <Table
        :title="'MOST ACTIVE THIS DAY'"
        :area="'activeDay'"
        :column="'Time'"
        :list="dailyActive"
        :format="identity"
      />

      <section style="grid-area:online">
        <h1>{{ playerCount }} PLAYERS ONLINE</h1>
        <table style="width:200px;margin: 0 auto;">
          <thead>
            <tr>
              <th>Username</th>
            </tr>
          </thead>

          <tbody>
            <tr v-for="(player, index) in online" v-if="index < onlineIndex" :key="index">
              <td>{{ player.Name }}</td>
            </tr>
            <tr v-if="online.length>4 && onlineIndex!=online.length">
              <td>...</td>
            </tr>
          </tbody>
        </table>
        <button
          @click="onlineIndex=online.length"
          v-if="online.length>4 && onlineIndex!=online.length"
        >Show all</button>
      </section>

      <Table
        :title="'MOST ACTIVE THIS WEEK'"
        :area="'activeWeek'"
        :column="'Time'"
        :list="weeklyActive"
        :format="identity"
      />
    </article>

    <!-- General stats -->
    <article id="general">
      <Table
        :title="'MOST SURVIVALS THIS WEEK'"
        :area="'survivedDay'"
        :column="'# Rounds'"
        :list="weeklySurvivals"
        :format="identity"
      />

      <Table
        :title="'MOST SURVIVALS THIS DAY'"
        :area="'survivedWeek'"
        :column="'# Rounds'"
        :list="dailySurvivals"
        :format="identity"
      />

      <Table
        :title="'MOST KILLS THIS DAY'"
        :area="'killsDay'"
        :column="'Kills'"
        :list="dailyKills"
        :format="identity"
      />

      <Table
        :title="'MOST KILLS THIS WEEK'"
        :area="'killsWeek'"
        :column="'Kills'"
        :list="weeklyKills"
        :format="identity"
      />
    </article>

    <!-- Current map -->
    <section style="grid-area:current">
      <h1>Current Map</h1>
      <img :src="mapImage" width="30%" height="30%">
    </section>

    <!-- Activity over 48 hours -->
    <article style="display: flex; justify-content: center;">
      <ActivityChart/>
    </article>

    <!-- human wins/zombie wins -->
    <article id="wins">
      <div style="grid-area: humans;">
        <h1>HUMAN WINS</h1>
        <h1>{{ humanWins }}</h1>
      </div>
      <div style="grid-area: zombies;">
        <h1>ZOMBIE WINS</h1>
        <h1>{{ zombieWins }}</h1>
      </div>
    </article>

    <!-- Pick/Kill distribution -->
    <article id="classes">
      <div style="grid-area: picks;">
        <h1>PICK DISTRIBUTION</h1>
        <DonutChart :series="chartPickDistribution"/>
      </div>
      <div style="grid-area: kills;">
        <h1>KILL DISTRIBUTION</h1>
        <DonutChart :series="chartKillDistribution"/>
      </div>
    </article>

    <!-- Survival/Score distribution -->
    <article id="class-survivals">
      <div style="grid-area: count;">
        <h1>CLASS MOST SURVIVED</h1>
        <DonutChart :series="classSurvivals"/>
      </div>
      <div style="grid-area: time;">
        <h1>CLASS SCORE DISTRIBUTION</h1>
        <DonutChart :series="chartScoreDistribution"/>
      </div>
    </article>

    <!-- Average alive time -->
    <div>
      <h1>CLASS AVERAGE ALIVE</h1>
      <SurvivalChart :series="chartAliveTime"/>
    </div>

    <!-- Class stats -->
    <h1>PLAYER'S MOST PICKED CLASS</h1>
    <article style="width:95%; margin: 0 auto;">
      <tabs :options="{ useUrlFragment: false }">
        <tab v-for="(role, index) in classes" :name="getName(role)" :key="index">
          <ClassPage :className="role"/>
        </tab>
      </tabs>
    </article>

    <!-- Top5 in class -->
    <article id="topfive">
      <List
          v-for="(role, index) in classes" :key="index"
          v-if="index < 5"
          :title="role"
          :area="role"
          :list="classBest(role)"
        />
      <List
        v-for="(role, index) in classes" :key="index"
        v-if="index >= 5"
        :title="role"
        :area="role"
        :list="classBest(role)"
      />
    </article>

    <!-- Map stats -->
    <h1>MAP STATISTICS</h1>
    <article>
      <MapStatistic v-for="(map, index) in maps" :key="index" :mapName="map"/>
    </article>

    <!-- Reset Counter -->
    <br>
    <br>
    <h1>Time until Reset</h1>
    <p>{{ resetTime }}</p>
    <br>
    <br>
  </main>
</template>

<script>
import moment from "moment";
import { mapGetters, mapState } from "vuex";

import Row from "@/components/Row.vue";
import Table from "@/components/Table.vue";
import List from "@/components/List.vue";
import ClassPage from "@/components/ClassPage.vue";
import MapStatistic from "@/components/MapStatistic.vue";
import ActivityChart from "@/components/ActivityChart.vue";
import DonutChart from "@/components/DonutChart.vue";
import SurvivalChart from "@/components/SurvivalChart.vue";

export default {
  name: "home",
  components: {
    Row,
    Table,
    List,
    ClassPage,
    MapStatistic,
    ActivityChart,
    DonutChart,
    SurvivalChart
  },
  data() {
    return {
      countdown: 0, // seconds left until reset

      identity: a => a,

      // amount of players shown
      onlineIndex: 4,

      maps: [
        "infc_warehouse",
        "infc_warehouse2",
        "infc_damascus",
        "infc_eidalfitr",
        "infc_newdust",
        "infc_halfdust",
        "infc_canyon",
        "infc_headquarter",
        "infc_hardcorepit",
        "infc_normandie_2k19",
        "infc_lunaroutpost",
        "infc_skull",
        "infc_toilet",
        "infc_toilet_old",
        "infc_malinalli_k9f",
        "infc_k9f_small",
        "infc_bamboo2",
        "infc_bamboo3",
        "infc_towers",
        "infc_spacelab",
        "infc_skull_winter",
        "infc_normandie_winter",
        //"infc_deathdealer",
        //'infc_cherrytemple',
        //'infc_crystalcave',
      ],

      classes: [
        "Soldier",
        "Scientist",
        "Biologist",
        "Medic",
        "Hero",
        "Ninja",
        "Mercenary",
        "Sniper",
        "Engineer",
        "Looper"
      ]
    };
  },
  beforeMount() {
    // Calculate next friday morning (00:00)
    const friday = 5;
    const today = moment.utc().isoWeekday();
    let nextReset;

    if (today < friday) {
      // then just give me this week's instance of that day
      nextReset = moment.utc().isoWeekday(friday);
    } else {
      // *next week's* instance of that same day
      nextReset = moment
        .utc()
        .add(1, "weeks")
        .isoWeekday(friday);
    }
    nextReset.utc().startOf("day");
    this.countdown = Math.floor(Math.abs(moment().diff(nextReset)) / 1000);

    setInterval(() => {
      this.countdown -= 1;
    }, 1000);
  },

  computed: {
    ...mapState([
      "mvp",
      "online",
      "zombieWins",
      "humanWins",
      "dailyActive",
      "weeklyActive",
      "dailySurvivals",
      "weeklySurvivals",
      "classSurvivals",
      "dailyKills",
      "weeklyKills"
    ]),
    ...mapGetters([
      "mapImage",
      "playerCount",
      "classBest",
      "chartPickDistribution",
      "chartKillDistribution",
      "chartScoreDistribution",
      "chartAliveTime"
    ]),
    resetTime() {
      let seconds = this.countdown;

      const days = Math.floor(seconds / (60 * 60 * 24));
      seconds -= days * (60 * 60 * 24);

      const hours = Math.floor(seconds / (60 * 60));
      seconds -= hours * (60 * 60);

      const minutes = Math.floor(seconds / 60);
      seconds -= minutes * 60;

      return `${days}d ${hours}h ${minutes}min ${seconds}s`;
    }
  },
  methods: {
    getName(role) {
      return `<img src='https://stats.resamvi.io/classes/${role}.png' width='75%' height='75%' style='margin: 0px auto'>`;
    }
  }
};
</script>

<style scoped>
#general {
  display: grid;
  grid-template-columns: 25% 25% 25% 25%;
  grid-template-rows: auto;
  grid-template-areas: "survivedDay survivedWeek killsDay killsWeek";
  margin-bottom: 5%;
}

#activity {
  display: grid;
  grid-template-columns: 33% 33% 33%;
  grid-template-rows: auto;
  grid-template-areas: "activeDay online activeWeek";
  margin-bottom: 5%;
}

#mvp {
  display: grid;
  grid-template-columns: 25% 50% 25%;
  grid-template-rows: auto;
  grid-template-areas: ".  mvp   .";
  margin-bottom: 5%;
}

#wins {
  display: grid;
  margin-bottom: 5%;
  grid-template-columns: 50% 50%;
  grid-template-rows: auto;
  grid-template-areas: "humans zombies";
}

#topfive {
  display: grid;
  grid-template-columns: 10% 16% 16% 16% 16% 16% 10%;
  grid-template-rows: auto;
  grid-template-areas: 
    ". Soldier Scientist Biologist Medic Hero ." 
    ". Ninja Mercenary Sniper Engineer Looper .";
}

#classes {
  display: grid;
  grid-template-columns: 50% 50%;
  grid-template-rows: auto;
  grid-template-areas: "picks kills";
}

#class-survivals {
  display: grid;
  grid-template-columns: 50% 50%;
  grid-template-rows: auto;
  grid-template-areas: "count time";
}

main {
  background-color: #f2f2f2;
  display: inline-block;
  width: 100%;
}

button {
  background: none;
  border: none;
  cursor: pointer;
}

.center {
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
