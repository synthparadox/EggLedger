<template>
    <input
        type="text" v-on:click="e => {e.preventDefault(); e.stopPropagation(); $emit('openModal', index, (orIndex ?? null));}"
        class="filter-select w-auto border-gray-300 text-gray-400" placeholder="Click to select" readonly
        :id="internalId" :value="modelValue">
</template>

<script>
    export default {
        emits: ['openModal'],
        errorCaptured(err, vm, info) {
            console.error('Error captured in component:', err);
            console.error('Vue instance:', vm);
            console.error('Error info:', info);
            // Return false to stop the error from propagating further to the global error handler
            return false;
        },
        props: {
            index: Number,
            orIndex: Number,
            modelValue: String,
            internalId: String,
        },
    }
</script>