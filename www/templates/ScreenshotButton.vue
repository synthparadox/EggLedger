<template>
    <button class="screenshot-button" v-on:click="takeScreenshot">
        <img src="/images/camera.png" alt="Take a screenshot">
    </button>
</template>

<script>
    export default {
        emits: ['take'],
        props: {
            elementId: String,
            schemaName: String,
        },
        methods: {
            takeScreenshot(){
                this.$emit('take', this.elementId, this.schemaName);
            }
        }
    }
</script>