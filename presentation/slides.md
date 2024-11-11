---
# You can also start simply with 'default'
theme: default
# random image from a curated Unsplash collection by Anthony
# like them? see https://unsplash.com/collections/94734566/slidev
background: /background.jpg
# some information about your slides (markdown enabled)
title: Разработка Cloud Native приложений на Go
info: |
  ## Slidev Starter Template
  Presentation slides for developers.

  Learn more at [Sli.dev](https://sli.dev)
# apply unocss classes to the current slide
class: text-center
# https://sli.dev/features/drawing
drawings:
  persist: false
# slide transition: https://sli.dev/guide/animations.html#slide-transitions
transition: slide-left
# enable MDC Syntax: https://sli.dev/features/mdc
mdc: true
---

# Разработка Cloud Native приложений на Go

<div class="abs-br m-6 flex gap-2">
  <a href="https://github.com/vokinneberg/ya-practicum-go-cloud-native-workshop" target="_blank" alt="GitHub" title="Open in GitHub"
    class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

---
layout: two-cols
layoutClass: gap-16
---

# Добро пожаловать на вебинар! 🤗

Пока мы ожидаем остальных, пожалуйста:

- Переименуйтесь в Zoom в формате Имя и Фамилия ✍️
- Если есть возможность, включите Web-камеру 📹

::right::

<Toc v-click minDepth="1" maxDepth="2"></Toc>

<style>

</style>

---
layout: image-right
image: https://memecreator.org/static/images/memes/5229661.jpg
---

# 🌥️ Что такое Cloud Native приложение и чем оно отличается от обычного приложения?

<v-click>

  - Почему Cloud Native?

</v-click>

<v-click>

  - Аттрибуты Cloud Native приложения

</v-click>

---
layout: image-right
image: https://memecreator.org/static/images/memes/5229661.jpg
---

# 🛠️ Cloud native паттерны в Go

<v-click>

  -  Circuit Breaker

</v-click>

<v-click>

  - Timeout

</v-click>

<v-click>

  - Sharding

</v-click>

---
layout: image-right
image: https://memecreator.org/static/images/memes/5229661.jpg
---

# ☸️ Основы работы приложения в Kubernetes

<v-click>

  - Стратегия развертывания

</v-click>

<v-click>

  - Настройка

</v-click>

<v-click>

  - Service discovery

</v-click>

<v-click>

  - Балансировка нагрузки и масштабирование

</v-click>

<v-click>

  - Телеметрияе

</v-click>

<v-click>

  - Обеспечение безопасности приложения

</v-click>

---
layout: image-right
image: https://memecreator.org/static/images/memes/5229661.jpg
---

# 🚀 Продвинутые паттерны развертывания приложений в Kubernetes

<v-click>

  - Sidecar

</v-click>

<v-click>

  - Elastic Scaling

</v-click>

---

# Вопросы?

---

# Источники

[Documentation](https://sli.dev) · [GitHub](https://github.com/slidevjs/slidev) · [Showcases](https://sli.dev/resources/showcases)

<PoweredBySlidev mt-10 />
