---
theme: default
background: /background.jpg
class: text-center
drawings:
  persist: false
transition: slide-left
mdc: true
title: Разработка Cloud Native приложений на Go
info: |
  <h2>Разработка Cloud Native приложений на Go</h2>
  <h5 class="font-light">Евгений Гребенников</h5>
---

<div class="text-center text-4xl">
  Разработка Cloud Native приложений на Go
</div>

<div class="text-center text-2xl">
  Евгений Гребенников
</div>

<div class="abs-br m-6 flex gap-2 items-center">
  <div class="text-center text-xl flex items-center gap-1">
    {{ new Date().toLocaleDateString('ru-RU', { year: 'numeric', month: 'long', day: 'numeric' }) }}
  </div>
  <a href="https://github.com/vokinneberg/ya-practicum-go-cloud-native-workshop" target="_blank" alt="GitHub" title="Open in GitHub"
    class="text-xl slidev-icon-btn opacity-50 !border-none !hover:text-white">
    <carbon-logo-github />
  </a>
</div>

---
layout: two-cols
hideInToc: true
---

## 🤗 Добро пожаловать на вебинар!

Пока мы ожидаем остальных, пожалуйста:

- ✍️ Переименуйтесь в Zoom в формате Имя и Фамилия
- 📹 Если есть возможность, включите Web-камеру

::right::

## План вебинара

<Toc minDepth="1" maxDepth="2" class="text-left" mode="all" />

---
layout: image-right
image: ./cloudnative-apps-what.jpg
title: 🌥️ Почему Cloud Native?
---

## 🌥️ Почему Cloud Native?

<v-click>
  <div class="bg-gray-200 p-4 rounded" style="margin: 20px;">
    <div class="text-left text-l">
      Cloud Native Is About Culture, Not Containers.
    </div>
    <div class="text-right text-sm italic">
      Cummins, Holly. Cloud Native London 2018.
    </div>
  </div>
</v-click>

<v-click>
  <div class="bg-gray-200 p-4 rounded" style="margin: 20px;">
    <div class="text-left text-l">
      Надежность (Dependability) компьютерной системы — это ее способность избегать сбоев, которые более часты или серьезны, а также периодов простоя, которые длиннее, чем это приемлемо для пользователей.
    </div>
    <div class="text-right text-sm italic">
      Фундаментальные концепции надежности компьютерных систем (2001).
    </div>
  </div>
</v-click>

---
layout: image-right
image: ./cloud-native-attrs.png
hideInToc: true
---

## 🧠 Атрибуты Cloud Native приложения

<v-click>

- **Availability** - Способность системы выполнять свою предназначенную функцию в случайный момент времени.

</v-click>

<v-click>

- **Reliability** - Способность системы выполнять свою предназначенную функцию в течение заданного интервала времени.

</v-click>

<v-click>

- **Maintainability** - Способность системы подвергаться модификациям и ремонту.

</v-click>

---
layout: image-right
class: text-center
image: ./12_factor_apps.png
hideInToc: true
---

<div class="flex justify-center items-center h-full text-2xl">
  <div>
    ⚡ The Twelve Factors
  </div>
</div>

---
layout: two-cols
title: ☸️ Основы разработки в Kubernetes
---

## ☸️ Основы разработки в Kubernetes

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

  - Телеметрия

</v-click>

<v-click>

  - Обеспечение безопасности приложения

</v-click>

::right::

<img src="./k8s_primitives.png" class="abs-tr m-3" width="50%">

---
layout: default
title: 🚀 Продвинутые паттерны развертывания приложений в Kubernetes
---

<div>
  <div class="flex justify-center items-center h-full">
    <div class="text-2xl">
      🚀 Продвинутые паттерны развертывания приложений в Kubernetes
    </div>
  </div>
  <div class="flex justify-center items-center h-full" style="margin-top: 0px;">
    <img src="./pattern.jpg" width="50%">
  </div>
</div>

---
layout: default
hideInToc: true
---

## 🚗 Sidecar

Sidecar — это архитектурный подход, при котором в одном поде (Pod) размещаются несколько контейнеров, выполняющих различные, но взаимодополняющие функции.

```mermaid
graph LR
    subgraph Pod
        direction LR
        A[Основное приложение]
        B[Sidecar контейнер]
    end
    A -- Взаимодействие --> B
    B -- Дополнительные функции --> A
```

---
layout: default
hideInToc: true
---

## 🐙 Service Mesh

Service Mesh — это инфраструктурный слой, предназначенный для управления взаимодействиями между микросервисами в распределенных системах.

```mermaid
graph LR
    subgraph Control Plane
        CP[Управление]
    end

    subgraph Pod1
        direction TB
        ServiceA[Сервис A]
        ProxyA[Sidecar Proxy]
    end

    subgraph Pod2
        direction TB
        ServiceB[Сервис B]
        ProxyB[Sidecar Proxy]
    end

    subgraph Pod3
        direction TB
        ServiceC[Сервис C]
        ProxyC[Sidecar Proxy]
    end

    ServiceA -- Локальное взаимодействие --> ProxyA
    ServiceB -- Локальное взаимодействие --> ProxyB
    ServiceC -- Локальное взаимодействие --> ProxyC

    ProxyA -- Трафик сервисов --> ProxyB
    ProxyB -- Трафик сервисов --> ProxyC
    ProxyC -- Трафик сервисов --> ProxyA

    CP -- Конфигурация и политики --> ProxyA
    CP -- Конфигурация и политики --> ProxyB
    CP -- Конфигурация и политики --> ProxyC
```

---
layout: default
hideInToc: true
---

## 🐙 Elastic Scaler

Elastic Scaler — это механизм автоматического масштабирования ресурсов, позволяющий динамически адаптировать количество подов или их ресурсы в зависимости от текущей нагрузки и требований приложения.

```mermaid
graph LR
    subgraph Kubernetes Cluster
     direction LR
        A[Пользователи] --> Service[Сервис] --> Pods[Pod1, Pod2, ..., PodN]

        Pods -- Метрики --> MetricsAPI[Метрики]

        HPA[Horizontal Pod Scaler] -- Запрос метрик --> MetricsAPI
        HPA --> |Масштабирование| Deployment[Deployment] --> Pods
    end
```

---
layout: center
title: 🤔 Вопросы?
---

 ## 🤔 Вопросы?

---
layout: center
title: 📚 Источники
---

## 📚 Источники

- [What is Cloud Native?](https://aws.amazon.com/what-is/cloud-native/)
- [Cloud Native Computing Foundation](https://www.cncf.io/)
- [The Twelve Factors](https://12factor.net/)
- [Red Hat Cloud Native Apps](https://www.redhat.com/en/topics/cloud-native-apps)
- [Cloud Native Go: Building Reliable Services in Unreliable Environments](https://a.co/d/gDIj5SJ)
- [Kubernetes Patterns: Reusable Elements for Designing Cloud Native](https://a.co/d/aRQ3F4X)

<PoweredBySlidev mt-10 />
