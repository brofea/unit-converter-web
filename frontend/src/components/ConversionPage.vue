<script setup lang="ts">
import { computed, ref } from 'vue'

interface ConvertRequest {
  value: number
  from: string
  to: string
}

interface ConvertResponse {
  result?: number
  error?: string
}

const props = defineProps<{
  category: string
  unitOptions: Array<{ label: string; value: string }>
  apiBaseUrl: string
}>()

const selectedUnitA = ref(props.unitOptions[0]?.value ?? '')
const selectedUnitB = ref(props.unitOptions[1]?.value ?? props.unitOptions[0]?.value ?? '')
const valueA = ref('')
const loading = ref(false)
const responseMessage = ref('')
const resultValue = ref<number | null>(null)

const endpoint = computed(() => `${props.apiBaseUrl}/api/v1/convert`)

const selectedFromLabel = computed(() => {
  return props.unitOptions.find((unit) => unit.value === selectedUnitA.value)?.label ?? selectedUnitA.value
})

const selectedToLabel = computed(() => {
  return props.unitOptions.find((unit) => unit.value === selectedUnitB.value)?.label ?? selectedUnitB.value
})

function swapUnits(): void {
  const temp = selectedUnitA.value
  selectedUnitA.value = selectedUnitB.value
  selectedUnitB.value = temp
}

async function submitConvert(): Promise<void> {
  if (!selectedUnitA.value || !selectedUnitB.value) {
    responseMessage.value = '请选择起始单位和目标单位。'
    return
  }

  if (valueA.value.trim() === '') {
    responseMessage.value = '请填写数值。'
    return
  }

  const value = Number(valueA.value)
  if (Number.isNaN(value)) {
    responseMessage.value = '请输入有效数值。'
    return
  }

  const payload: ConvertRequest = {
    value,
    from: selectedUnitA.value,
    to: selectedUnitB.value
  }

  loading.value = true
  responseMessage.value = ''
  resultValue.value = null

  try {
    const response = await fetch(endpoint.value, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    })

    if (!response.ok) {
      const errorData = (await response.json().catch(() => null)) as ConvertResponse | null
      throw new Error(errorData?.error ?? `请求失败：${response.status}`)
    }

    const data = (await response.json()) as ConvertResponse

    if (typeof data.result !== 'number') {
      throw new Error(data.error ?? '后端未返回有效结果。')
    }

    resultValue.value = data.result
    responseMessage.value = `转换成功：${value} ${selectedFromLabel.value} = ${data.result} ${selectedToLabel.value}`
  } catch (error) {
    const message = error instanceof Error ? error.message : '未知错误'
    responseMessage.value = `转换失败：${message}`
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <section class="panel">
    <h2 class="title">{{ category }}转换</h2>

    <div class="form-row">
      <label for="fromUnit" class="label">起始单位</label>
      <select id="fromUnit" v-model="selectedUnitA" class="input">
        <option v-for="unit in unitOptions" :key="unit.value" :value="unit.value">
          {{ unit.label }}（{{ unit.value }}）
        </option>
      </select>
    </div>

    <div class="form-row">
      <label for="valueA" class="label">数值</label>
      <input id="valueA" v-model="valueA" type="text" class="input" placeholder="请输入数值" inputmode="decimal" />
    </div>

    <div class="form-row">
      <label for="toUnit" class="label">目标单位</label>
      <select id="toUnit" v-model="selectedUnitB" class="input">
        <option v-for="unit in unitOptions" :key="`to-${unit.value}`" :value="unit.value">
          {{ unit.label }}（{{ unit.value }}）
        </option>
      </select>
    </div>

    <div class="form-row">
      <label for="resultBox" class="label">结果</label>
      <input
        id="resultBox"
        class="input"
        type="text"
        :value="resultValue === null ? '' : String(resultValue)"
        placeholder="转换结果会显示在这里"
        readonly
      />
    </div>

    <p v-if="responseMessage" class="meta">{{ responseMessage }}</p>

    <div class="action-row">
      <button class="swap-btn" type="button" @click="swapUnits">调换单位</button>

      <button class="convert-btn" type="button" :disabled="loading" @click="submitConvert">
        {{ loading ? '转换中...' : '开始转换' }}
      </button>
    </div>
  </section>
</template>

<style scoped>
.panel {
  width: 100%;
  max-width: 860px;
  box-sizing: border-box;
  margin: 0 auto;
  overflow-x: clip;
  padding: 1.6rem;
  border: 1px solid rgba(19, 28, 45, 0.2);
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.88);
  box-shadow: 0 14px 36px rgba(18, 41, 76, 0.12);
  backdrop-filter: blur(8px);
}

.title {
  margin: 0 0 1.2rem;
  font-size: clamp(1.5rem, 1.4rem + 0.5vw, 1.9rem);
  letter-spacing: 0.03em;
}

.form-row {
  display: grid;
  min-width: 0;
  gap: 0.6rem;
  margin-bottom: 1rem;
}

.label {
  color: #20334f;
  font-weight: 600;
}

.input {
  display: block;
  width: 100%;
  max-width: 100%;
  box-sizing: border-box;
  padding: 0.75rem 0.85rem;
  border-radius: 12px;
  border: 1px solid #91a7c5;
  background: #fcfdff;
  color: #122034;
  font-size: 1rem;
}

select.input {
  appearance: none;
  background-image:
    linear-gradient(45deg, transparent 50%, #1b3a60 50%),
    linear-gradient(135deg, #1b3a60 50%, transparent 50%);
  background-position:
    calc(100% - 18px) calc(50% - 3px),
    calc(100% - 12px) calc(50% - 3px);
  background-size: 6px 6px, 6px 6px;
  background-repeat: no-repeat;
  padding-right: 2.1rem;
}

.input:focus {
  outline: 2px solid #1f6feb;
  outline-offset: 1px;
}

.action-row {
  display: flex;
  flex-wrap: wrap;
  gap: 0.7rem;
  margin-top: 0.3rem;
}

.swap-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.75rem 1rem;
  border: 1px solid #597da8;
  border-radius: 12px;
  background: #f3f8ff;
  color: #1b3a60;
  font-weight: 700;
  cursor: pointer;
}

.swap-btn:hover {
  background: #e5f0ff;
}

.convert-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.75rem 1.2rem;
  border: 0;
  border-radius: 12px;
  background: linear-gradient(120deg, #1259c3, #2880e5);
  color: #fff;
  font-weight: 700;
  letter-spacing: 0.02em;
  cursor: pointer;
  transition: transform 120ms ease, box-shadow 120ms ease;
}

.convert-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 10px 24px rgba(18, 89, 195, 0.35);
}

.convert-btn:disabled {
  opacity: 0.7;
  cursor: wait;
  transform: none;
  box-shadow: none;
}

.meta {
  margin: 0.75rem 0 0;
  color: #1d3150;
}

@media (max-width: 640px) {
  .panel {
    padding: 1.1rem;
    border-radius: 16px;
  }
}
</style>
