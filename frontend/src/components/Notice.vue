<template>
  <v-snackbar right top v-model="snackbar" :timeout="timeout" :color="color">
    {{ text }}
    <v-btn color="secondary" icon @click="snackbar = false">
      <v-icon>mdi-close</v-icon>
    </v-btn>
  </v-snackbar>
</template>

<script>
export default {
  data: () => ({
    snackbar: false,
    color: "",
    text: "",
    timeout: 3000
  }),
  methods: {
    show(text, c) {
      this.text = text;
      this.color = c == 0 ? "black" : "error";
      this.snackbar = true;
    }
  }
};
</script>
