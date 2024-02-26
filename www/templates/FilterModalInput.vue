<template>
    <input
        type="text" v-on:click="e => {e.preventDefault(); e.stopPropagation(); $emit('openModal', index, (orIndex ?? null));}"
        class="filter-select w-auto border-gray-300 text-gray-400" placeholder="Click to select" readonly
        :id="internalId" :value="modelValue">
</template>

<script>
    export default {
        emits: ['openModal'],
        props: {
            index: Number,
            orIndex: Number,
            modelValue: String,
            internalId: String,
        },
    }
</script>