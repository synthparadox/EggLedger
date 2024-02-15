<template>
    <select class="filter-select border-gray-300 text-gray-400"
        v-on:change="handleFilterChange"
        :id="internalId"
        :value="modelValue"
    >
        <option 
            v-for="option in optionList" 
            :value="option.value" 
            :key="option"
            :class="'filter-select-option ' + (level == 'value' ? (option.styleClass ?? '') : '')"
        >
            {{ option.text }}
        </option>
    </select>
</template>

<script>
    export default {
        emits: ['changeFilterValue', 'handleFilterChange'],
        errorCaptured(err, vm, info) {
            console.error('Error captured in component:', err);
            console.error('Vue instance:', vm);
            console.error('Error info:', info);
            // Return false to stop the error from propagating further to the global error handler
            return false;
        },
        props: {
            internalId: String,
            optionList: Array,
            level: String,
            index: Number,
            orIndex: Number,
            modelValue: String,
        },
        methods: {
            handleFilterChange(event){
                this.$emit('changeFilterValue', this.index, this.orIndex, this.level, event.target.value);
                this.$emit('handleFilterChange', event);
            }
        }
    }
</script>