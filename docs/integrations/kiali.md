# Kiali

## Use Kiali with iter8
Kiali is an observability console for Istio with service mesh configuration capabilities. It helps you to understand the structure of your service mesh by inferring the topology, and also provides the health of your mesh. Kiali provides detailed metrics, and a basic Grafana integration is available for advanced queries. Distributed tracing is provided by integrating Jaeger. For detailed installation and information about Kiali in general, please reference [Kiali.io](https://kiali.io)

## Enabling the iter8 Console in Kiali

Currently the iter8 extension for Kiali (v1.24) works only for [iter8 v1.0.0](https://v1-0-0.iter8.tools/) and above. To install on either Kubernetes or OpenShift, please see [here](https://v1-0-0.iter8.tools/installation/).

### Install Kiali Using Operator

Kiali version 1.18.1 (and above) provides the capability to observe iter8 experiment runtime behavior. To enable the iter8 extensions for Kiali, which are disabled by default, you must update the Kiali CR using the Kiali operator.

To check if Kiali operator is installed, use:

```bash
kubectl --namespace kiali-operator get pods
```

To install the Kiali operator, follow the steps in [Install Kiali Latest]( https://kiali.io/documentation/latest/installation-guide/#_install_kiali_latest). You can verify that the Kiali CR is created by using the command:

```bash
kubectl  --namespace istio-system get kialis.kiali.io kiali
```

If this is the new installation, you will be asked to choose an authentication strategy (`login`, `anonymous`, `ldap`, `openshift`, `token` or `openid`). Depending on the chosen strategy, the installation process may prompt for additional information. Please see [Kiali Login Options](https://kiali.io/documentation/latest/installation-guide/#_login_options) for details about authentication strategies.

### Enable iter8 in the Kiali Operator CR

Follow the step [Create or Edit the Kiali CR](https://kiali.io/documentation/latest/installation-guide/#_create_or_edit_the_kiali_cr) or use:

```bash
kubectl --namespace istio-system edit kialis.kiali.io kiali
```

Find the `iter_8` key under `spec.extensions` and set `enabled` to `true`. The relevant portion of the CR is:

```
# Kiali enabled integration with iter8 project.
# If this extension is enabled, Kiali will communicate with iter8 controller allowing to manage Experiments and review results.
# Additional documentation https://v1-0-0.iter8.tools/
#    ---
#    iter_8:
#
# Flag to indicate if iter8 extension is enabled in Kiali
#      ---
#      enabled: false
#
extensions:
  iter_8:
    enabled: true
```

Changing the Kialis.kiali.io CR will trigger Kiali to re-configured. You can also force the Kiali pod to restart.
Restart the Kiali pods:

```bash
kubectl --namespace istio-system delete pod $(kubectl --namespace istio-system get pod --selector='app=kiali' -o jsonpath='{.items[0].metadata.name}')
```

You can check if the pod has successfully restarted by inspecting the pods:

```bash
kubectl --namespace istio-system get pods
```

Install iter8 v1.0.0-rc3. See [install instructions](https://v1-0-0.iter8.tools/installation)

Start Kiali using:

```bash
istioctl dashboard kiali
```

## Features of the iter8 Extension for Kiali

### Experiments Overview

{{< figure src="/images/kiali-list-experiments.png" title="iter8 main page" caption="Iter8 main page lists all the experiments in available namespace(s).">}}

### Create Experiment (Form Style)

You can create new experiment from the Action pulldown on the right of the listing page.

{{< figure src="/images/kiali-experiment-create-1.png" title="Experiment creation">}}

{{< figure src="/images/kiali-experiment-create-2.png" title="Experiment creation -- additional configuration options">}}

{{< figure src="/images/kiali-experiment-create-duration-traffic.png" title="Experiment creation -- Duration and Traffic options">}}

{{< figure src="/images/kiali-experiment-create-networking.png" title="Experiment creation -- Networking options">}}

### Create Experiment from File

You can create experiment from exiting local file, import for Github or manual enter the YAML in the YAML editor.
{{< figure src="/images/kiali-experiment-create-from-file.png" title="Experiment creation from file">}}

### Experiment Detail

{{< figure src="/images/kiali-experiment-detail.png" title="Experiment detail page" caption="Click on the name of the experiment from the listing page will show the experiment detail page. In the detail page, user can `pause`, `resume`, `terminate with traffic split` from the action pulldown. User can also `delete` experiment.">}}

{{< figure src="/images/kiali-experiment-detail-assessment.png" title="Experiment assessment page" caption="Click on the assessment tab will show the experiment assessment information.">}}

### Short Video about Kiali and Iter8 Integration

 {{< youtube bGEJLPHUZiQ>}}

## Troubleshooting Guide

**Issue**: Cannot find the Kiali CR in namespace `kiali-operator`.

Try using this command to install and start operator

```bash
bash <(curl -L https://kiali.io/getLatestKialiOperator) --accessible-namespaces '**' -oiv latest -kiv latest --operator-install-kiali true
```

---

**Issue**: The iter8 extension is not visible in Kiali

Check the configmap `kiali` using this command:

 ```bash
 kubectl  --namespace istio-systemedit configmap kiali
 ```

Ensure that `spec.extensions.iter_8.enabled` is set to `true`. To ensure that this configuration has taken effect, restart the Kiali pod:

```bash
kubectl --namespace istio-system delete pod $(kubectl --namespace istio-system get pod --selector='app=kiali' -o jsonpath='{.items[0].metadata.name}')
```

---

**Issue**: Error message `Kiali has iter8 extension enabled but it is not detected in the cluster`

Make sure iter8 is installed, check that both iter8-controller and iter8-analytics are functioning:

```bash
kubectl --namespace iter8 get pods
```

If you started Kiali before you installed iter8, please restart Kiali in the `istio-system` namespace to re-detect the iter8 installation.

---

**Issue**: Experiment(s) are missing in the iter8 main page

Make sure the namespace that contains the experiment is included in the Kiali accessible namespace `accessible_namespaces:` definitions in the CR.