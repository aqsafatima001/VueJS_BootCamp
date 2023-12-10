<template>
  <div class="hello">
    <div>
      <h1 class="title">Calculator</h1>
    </div>
    <hr>
    <div>
      <b-container fluid>
        <b-row class="text-center">
          <b-col cols="2" class="text-center"></b-col>
          <b-col cols="4" class="text-center">
            <form @submit.prevent="calculate">
              <div class="field">
                <label class="label">First Number</label>
                <b-form-input name="num1" v-model="num1" type="text"></b-form-input>
              </div>

              <div class="field" style="margin-top:13px;">
                <label class="label">Second Number</label>
                <b-form-input name="num2" v-model="num2" type="text"></b-form-input>
              </div>
              <b-button variant="primary" type="submit">Calculate</b-button>
            </form>
          </b-col>

          <b-col cols="4" class="text-left">
            <div><label class="label">Addition: {{ add }}</label></div>
            <div><label class="label">Multiplication: {{ mul }}</label></div>
            <div><label class="label">Subtraction: {{ sub }}</label></div>
            <div><label class="label">Division: {{ div }}</label></div>
          </b-col>
        </b-row>
      </b-container>
    </div>

    <hr>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Calculator',
  
  data: function() {
    return {
      add: "",
      mul: "",
      sub: "",
      div: "",
      num1: "",
      num2: ""
    }
  },
  methods: {
    calculate: function() {
      var data = { "num1": parseFloat(this.num1), "num2": parseFloat(this.num2) };
      axios.post("http://localhost:8090/calc", data, { headers: { "content-type": "text/plain" } })
        .then(result => {
          this.add = result.data['add'];
          this.mul = result.data['mul'];
          this.sub = result.data['sub'];
          this.div = result.data['div'];
        })
        .catch(error => {
          console.error(error);
        });
    }
  }
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
