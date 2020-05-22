<template>
  <v-row>
    <v-col cols="12">
      <v-card max-width="500" class="mx-auto">
        <v-list-item>
          <v-list-item-avatar>
            <img src="@/assets/logo.jpg" alt="Logo" />
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title class="headline">华工机器人实验室巨献</v-list-item-title>
            <v-list-item-subtitle>SCUT Robot Lab Present</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-card-text></v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn icon href="https://github.com/scutrobotlab/Robot_Monitor_Web">
            <v-icon>mdi-github</v-icon>
          </v-btn>
          <v-btn icon v-on:click="hc = 'pink'" v-bind:color="hc">
            <v-icon>mdi-heart</v-icon>
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  data: () => ({
    hc: ""
  })
};
</script>
