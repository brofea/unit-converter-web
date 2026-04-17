<script setup lang="ts">
import { computed, ref } from 'vue'

interface ConvertRequest {
  category: string
  fromUnit: string
  toUnit: string
  value: number
}

interface ConvertResponse {
  status: string
  result?: number
}

const props = defineProps<{
  category: string
  unitOptions: string[]
  apiBaseUrl: string
}>()

const selectedUnitA = ref(props.unitOptions[0] ?? '')
const selectedUnitB = ref(props.unitOptions[1] ?? props.unitOptions[0] ?? '')
const valueA = ref<number | null>(null)
const loading = ref(false)
const responseMessage = ref('')
const resultValue = ref<number | null>(null)

const endpoint = computed(() => `${props.apiBaseUrl}/convert`)

function pickUnitA(unit: string): void {
  selectedUnitA.value = unit
}

function pickUnitB(unit: string): void {
  selectedUnitB.value = unit
}

function swapUnits(): void {
  const temp = selectedUnitA.value
  selectedUnitA.value = selectedUnitB.value
  selectedUnitB.value = temp
}

async function submitConvert(): Promise<void> {
  if (!selectedUnitA.value || !selectedUnitB.value) {
    responseMessage.value = '请选择 A 单位和 B 单位。'
    return
  }

  if (valueA.value === null) {
    responseMessage.value = '请填写 A 数值。'
    return
  }

  const payload: ConvertRequest = {
    category: props.category,
    fromUnit: selectedUnitA.value,
    toUnit: selectedUnitB.value,
    value: Number(valueA.value)
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
      throw new Error(`请求失败：${response.status}`)
    }

    const data = (await response.json()) as ConvertResponse

    if (data.status !== 'success') {
      throw new Error('后端返回失败状态。')
    }

    resultValue.value = typeof data.result === 'number' ? data.result : null
    responseMessage.value = '转换成功。'
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

    <div class="form-row units">
      <p class="label">A 单位</p>
      <div class="unit-list">
        <label v-for="unit in unitOptions" :key="unit" class="unit-item">
          <input
            type="checkbox"
            :checked="selectedUnitA === unit"
            @change="pickUnitA(unit)"
          />
          <span>{{ unit }}</span>
        </label>
      </div>
    </div>

    <div class="form-row">
      <label for="valueA" class="label">A 数值</label>
      <input id="valueA" v-model.number="valueA" type="number" class="input" placeholder="请输入 A 数值" />
    </div>

    <div class="form-row units">
      <p class="label">B 单位</p>
      <div class="unit-list">
        <label v-for="unit in unitOptions" :key="`b-${unit}`" class="unit-item">
          <input
            type="checkbox"
            :checked="selectedUnitB === unit"
            @change="pickUnitB(unit)"
          />
          <span>{{ unit }}</span>
        </label>
      </div>
    </div>

    <div class="form-row">
      <label for="resultBox" class="label">B 数值</label>
      <input
        id="resultBox"
        class="input"
        type="text"
        :value="resultValue === null ? '' : String(resultValue)"
        placeholder="转换结果会显示在这里"
        readonly
      />
    </div>

    <div class="action-row">
      <button class="swap-btn" type="button" @click="swapUnits">调换单位</button>

      <button class="convert-btn" :disabled="loading" @click="submitConvert">
        {{ loading ? '转换中...' : '转换按钮' }}
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

.units {
  margin-bottom: 1.2rem;
}

.label {
  color: #20334f;
  font-weight: 600;
}

.unit-list {
  display: flex;
  flex-wrap: wrap;
  gap: 0.7rem;
}

.unit-item {
  display: inline-flex;
  align-items: center;
  gap: 0.4rem;
  padding: 0.45rem 0.65rem;
  border-radius: 999px;
  border: 1px solid #7e96b7;
  background: #f2f6fc;
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

.meta,
.status {
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
