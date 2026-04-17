# 项目指引：Unit Coverter

一个前后端分离的单位转换器项目

## 技术栈
- 前端：Vue 3, Vite
- 后端：Go, Gin

## 代码规范
- 前端：严禁使用 Options API，必须使用 <script setup> 语法。
- 前端：UI 组件优先使用原生的 HTML5 输入框，保持简洁。
- 后端：所有的 API 响应必须符合 JSON 格式：{ "status": "success", "result": float }。

## 业务逻辑
- 单位转换的核心逻辑应放在后端，前端仅负责发送请求和展示。
