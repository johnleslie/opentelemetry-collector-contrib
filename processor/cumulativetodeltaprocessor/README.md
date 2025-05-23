# Cumulative to Delta Processor
<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [beta]: metrics   |
| Distributions | [contrib], [k8s] |
| Warnings      | [Statefulness](#warnings) |
| Issues        | [![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Aprocessor%2Fcumulativetodelta%20&label=open&color=orange&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aopen+is%3Aissue+label%3Aprocessor%2Fcumulativetodelta) [![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Aprocessor%2Fcumulativetodelta%20&label=closed&color=blue&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aclosed+is%3Aissue+label%3Aprocessor%2Fcumulativetodelta) |
| Code coverage | [![codecov](https://codecov.io/github/open-telemetry/opentelemetry-collector-contrib/graph/main/badge.svg?component=processor_cumulativetodelta)](https://app.codecov.io/gh/open-telemetry/opentelemetry-collector-contrib/tree/main/?components%5B0%5D=processor_cumulativetodelta&displayType=list) |
| [Code Owners](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/CONTRIBUTING.md#becoming-a-code-owner)    | [@TylerHelmuth](https://www.github.com/TylerHelmuth) |

[beta]: https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/component-stability.md#beta
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
[k8s]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-k8s
<!-- end autogenerated section -->

## Description

The cumulative to delta processor (`cumulativetodeltaprocessor`) converts monotonic, cumulative sum and histogram metrics to monotonic, delta metrics. Non-monotonic sums and exponential histograms are excluded.

## Configuration

Configuration is specified through a list of metrics. The processor uses metric names to identify a set of cumulative metrics and converts them from cumulative to delta.

The following settings can be optionally configured:

- `include`: List of metrics names (case-insensitive), patterns or metric types to convert to delta. Valid values are: `sum`, `histogram`.
- `exclude`: List of metrics names (case-insensitive), patterns or metric types to not convert to delta.  **If a metric name matches both include and exclude, exclude takes precedence.** Valid values are: `sum`, `histogram`.
- `max_staleness`: The total time a state entry will live past the time it was last seen. Set to 0 to retain state indefinitely. Default: 0
- `initial_value`: Handling of the first observed point for a given metric identity.
  When the collector (re)starts, there's no record of how much of a given cumulative counter has already been converted to delta values.
  - `auto` (default): Send if and only if the startime is set AND the starttime happens after the component started AND the starttime is different from the timestamp.
    Suitable for gateway deployments, this heuristic is like `drop`, but keeps values for newly started counters (which could not have had previous observed values).
  - `keep`: Send the observed value as the delta value.
    Suitable for when the incoming metrics have not been observed before,
    e.g. running the collector as a sidecar, the collector lifecycle is tied to the metric source.
  - `drop`: Keep the observed value but don't send.
    Suitable for gateway deployments, guarantees that all delta counts it produces haven't been observed before, but loses the values between thir first 2 observations.

If neither include nor exclude are supplied, no filtering is applied.

#### Examples

```yaml
processors:
    # processor name: cumulativetodelta
    cumulativetodelta:

        # list the exact cumulative sum or histogram metrics to convert to delta
        include:
            metrics:
                - <metric_1_name>
                - <metric_2_name>
                .
                .
                - <metric_n_name>
            match_type: strict
```

```yaml
processors:
    # processor name: cumulativetodelta
    cumulativetodelta:

        # Convert all sum metrics
        include:
            metric_types:
              - sum
```

```yaml
processors:
    # processor name: cumulativetodelta
    cumulativetodelta:

        # Convert cumulative sum or histogram metrics to delta
        # if and only if 'metric' is in the name
        include:
            metrics:
                - ".*metric.*"
            match_type: regexp
```

```yaml
processors:
    # processor name: cumulativetodelta
    cumulativetodelta:

        # Convert cumulative sum metrics to delta
        # if and only if 'metric' is in the name
        include:
            metrics:
                - ".*metric.*"
            match_type: regexp
            metric_types:
              - sum
```

```yaml
processors:
    # processor name: cumulativetodelta
    cumulativetodelta:

        # Convert cumulative sum or histogram metrics to delta
        # if and only if 'metric' is not in the name
        exclude:
            metrics:
                - ".*metric.*"
            match_type: regexp
```

```yaml
processors:
    # processor name: cumulativetodelta
    cumulativetodelta:

        # Convert cumulative sum metrics with 'metric' in their name,
        # but exclude histogram metrics
        include:
            metrics:
                - ".*metric.*"
            match_type: regexp
        exclude:
          metric_types:
            - histogram
```

```yaml
processors:
    # processor name: cumulativetodelta
    cumulativetodelta:
        # If include/exclude are not specified
        # convert all cumulative sum or histogram metrics to delta
```

## Warnings

- [Statefulness](https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/standard-warnings.md#statefulness): The cumulativetodelta processor's calculates delta by remembering the previous value of a metric.  For this reason, the calculation is only accurate if the metric is continuously sent to the same instance of the collector.  As a result, the cumulativetodelta processor may not work as expected if used in a deployment of multiple collectors.  When using this processor it is best for the data source to being sending data to a single collector.


[beta]: https://github.com/open-telemetry/opentelemetry-collector#beta
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
