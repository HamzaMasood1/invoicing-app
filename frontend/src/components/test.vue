<script lang="ts" setup>
import {reactive} from "vue";
import {Test} from "../../wailsjs/go/main/App";

const data = reactive({
    name: "hellooooo",
})

function test(){
    Test("wow").then(result =>{
        data.name = result
    })
}
</script>


<template>
  <p>{{data.name}}</p>
  <button @click="test">click me~!</button>
</template>