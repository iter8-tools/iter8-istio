// +build !ignore_autogenerated

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

package v1alpha2

import (
	apiv1alpha2 "github.com/iter8-tools/iter8/pkg/analytics/api/v1alpha2"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Assessment) DeepCopyInto(out *Assessment) {
	*out = *in
	in.Baseline.DeepCopyInto(&out.Baseline)
	if in.Candidates != nil {
		in, out := &in.Candidates, &out.Candidates
		*out = make([]VersionAssessment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Winner != nil {
		in, out := &in.Winner, &out.Winner
		*out = new(WinnerAssessment)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Assessment.
func (in *Assessment) DeepCopy() *Assessment {
	if in == nil {
		return nil
	}
	out := new(Assessment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ExperimentCondition)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CounterMetric) DeepCopyInto(out *CounterMetric) {
	*out = *in
	if in.PreferredDirection != nil {
		in, out := &in.PreferredDirection, &out.PreferredDirection
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CounterMetric.
func (in *CounterMetric) DeepCopy() *CounterMetric {
	if in == nil {
		return nil
	}
	out := new(CounterMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Criterion) DeepCopyInto(out *Criterion) {
	*out = *in
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(Threshold)
		(*in).DeepCopyInto(*out)
	}
	if in.IsReward != nil {
		in, out := &in.IsReward, &out.IsReward
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Criterion.
func (in *Criterion) DeepCopy() *Criterion {
	if in == nil {
		return nil
	}
	out := new(Criterion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Duration) DeepCopyInto(out *Duration) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.MaxIterations != nil {
		in, out := &in.MaxIterations, &out.MaxIterations
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Duration.
func (in *Duration) DeepCopy() *Duration {
	if in == nil {
		return nil
	}
	out := new(Duration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Experiment) DeepCopyInto(out *Experiment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Experiment.
func (in *Experiment) DeepCopy() *Experiment {
	if in == nil {
		return nil
	}
	out := new(Experiment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Experiment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentCondition) DeepCopyInto(out *ExperimentCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentCondition.
func (in *ExperimentCondition) DeepCopy() *ExperimentCondition {
	if in == nil {
		return nil
	}
	out := new(ExperimentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentList) DeepCopyInto(out *ExperimentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Experiment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentList.
func (in *ExperimentList) DeepCopy() *ExperimentList {
	if in == nil {
		return nil
	}
	out := new(ExperimentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExperimentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentSpec) DeepCopyInto(out *ExperimentSpec) {
	*out = *in
	in.Service.DeepCopyInto(&out.Service)
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]Criterion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrafficControl != nil {
		in, out := &in.TrafficControl, &out.TrafficControl
		*out = new(TrafficControl)
		(*in).DeepCopyInto(*out)
	}
	if in.AnalyticsEndpoint != nil {
		in, out := &in.AnalyticsEndpoint, &out.AnalyticsEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(Duration)
		(*in).DeepCopyInto(*out)
	}
	if in.Cleanup != nil {
		in, out := &in.Cleanup, &out.Cleanup
		*out = new(bool)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(Metrics)
		(*in).DeepCopyInto(*out)
	}
	if in.ManualOverride != nil {
		in, out := &in.ManualOverride, &out.ManualOverride
		*out = new(ManualOverride)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentSpec.
func (in *ExperimentSpec) DeepCopy() *ExperimentSpec {
	if in == nil {
		return nil
	}
	out := new(ExperimentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentStatus) DeepCopyInto(out *ExperimentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ExperimentCondition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.InitTimestamp != nil {
		in, out := &in.InitTimestamp, &out.InitTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.EndTimestamp != nil {
		in, out := &in.EndTimestamp, &out.EndTimestamp
		*out = (*in).DeepCopy()
	}
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
	if in.CurrentIteration != nil {
		in, out := &in.CurrentIteration, &out.CurrentIteration
		*out = new(int32)
		**out = **in
	}
	if in.GrafanaURL != nil {
		in, out := &in.GrafanaURL, &out.GrafanaURL
		*out = new(string)
		**out = **in
	}
	if in.Assessment != nil {
		in, out := &in.Assessment, &out.Assessment
		*out = new(Assessment)
		(*in).DeepCopyInto(*out)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.AnalysisState != nil {
		in, out := &in.AnalysisState, &out.AnalysisState
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.EffectiveHosts != nil {
		in, out := &in.EffectiveHosts, &out.EffectiveHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentStatus.
func (in *ExperimentStatus) DeepCopy() *ExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(ExperimentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPMatchRequest) DeepCopyInto(out *HTTPMatchRequest) {
	*out = *in
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(StringMatch)
		(*in).DeepCopyInto(*out)
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(StringMatch)
		(*in).DeepCopyInto(*out)
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(StringMatch)
		(*in).DeepCopyInto(*out)
	}
	if in.Authority != nil {
		in, out := &in.Authority, &out.Authority
		*out = new(StringMatch)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]StringMatch, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.SourceLabels != nil {
		in, out := &in.SourceLabels, &out.SourceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.QueryParams != nil {
		in, out := &in.QueryParams, &out.QueryParams
		*out = make(map[string]StringMatch, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPMatchRequest.
func (in *HTTPMatchRequest) DeepCopy() *HTTPMatchRequest {
	if in == nil {
		return nil
	}
	out := new(HTTPMatchRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Host) DeepCopyInto(out *Host) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Host.
func (in *Host) DeepCopy() *Host {
	if in == nil {
		return nil
	}
	out := new(Host)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualOverride) DeepCopyInto(out *ManualOverride) {
	*out = *in
	if in.TrafficSplit != nil {
		in, out := &in.TrafficSplit, &out.TrafficSplit
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualOverride.
func (in *ManualOverride) DeepCopy() *ManualOverride {
	if in == nil {
		return nil
	}
	out := new(ManualOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Match) DeepCopyInto(out *Match) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]*HTTPMatchRequest, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HTTPMatchRequest)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Match.
func (in *Match) DeepCopy() *Match {
	if in == nil {
		return nil
	}
	out := new(Match)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metrics) DeepCopyInto(out *Metrics) {
	*out = *in
	if in.CounterMetrics != nil {
		in, out := &in.CounterMetrics, &out.CounterMetrics
		*out = make([]CounterMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RatioMetrics != nil {
		in, out := &in.RatioMetrics, &out.RatioMetrics
		*out = make([]RatioMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metrics.
func (in *Metrics) DeepCopy() *Metrics {
	if in == nil {
		return nil
	}
	out := new(Metrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RatioMetric) DeepCopyInto(out *RatioMetric) {
	*out = *in
	if in.ZeroToOne != nil {
		in, out := &in.ZeroToOne, &out.ZeroToOne
		*out = new(bool)
		**out = **in
	}
	if in.PreferredDirection != nil {
		in, out := &in.PreferredDirection, &out.PreferredDirection
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RatioMetric.
func (in *RatioMetric) DeepCopy() *RatioMetric {
	if in == nil {
		return nil
	}
	out := new(RatioMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.ObjectReference != nil {
		in, out := &in.ObjectReference, &out.ObjectReference
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.Candidates != nil {
		in, out := &in.Candidates, &out.Candidates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]Host, len(*in))
		copy(*out, *in)
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringMatch) DeepCopyInto(out *StringMatch) {
	*out = *in
	if in.Exact != nil {
		in, out := &in.Exact, &out.Exact
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringMatch.
func (in *StringMatch) DeepCopy() *StringMatch {
	if in == nil {
		return nil
	}
	out := new(StringMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Threshold) DeepCopyInto(out *Threshold) {
	*out = *in
	if in.CutoffTrafficOnViolation != nil {
		in, out := &in.CutoffTrafficOnViolation, &out.CutoffTrafficOnViolation
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Threshold.
func (in *Threshold) DeepCopy() *Threshold {
	if in == nil {
		return nil
	}
	out := new(Threshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficControl) DeepCopyInto(out *TrafficControl) {
	*out = *in
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(StrategyType)
		**out = **in
	}
	if in.OnTermination != nil {
		in, out := &in.OnTermination, &out.OnTermination
		*out = new(OnTerminationType)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(Match)
		(*in).DeepCopyInto(*out)
	}
	if in.Percentage != nil {
		in, out := &in.Percentage, &out.Percentage
		*out = new(int32)
		**out = **in
	}
	if in.MaxIncrement != nil {
		in, out := &in.MaxIncrement, &out.MaxIncrement
		*out = new(int32)
		**out = **in
	}
	if in.RouterID != nil {
		in, out := &in.RouterID, &out.RouterID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficControl.
func (in *TrafficControl) DeepCopy() *TrafficControl {
	if in == nil {
		return nil
	}
	out := new(TrafficControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionAssessment) DeepCopyInto(out *VersionAssessment) {
	*out = *in
	in.VersionAssessment.DeepCopyInto(&out.VersionAssessment)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionAssessment.
func (in *VersionAssessment) DeepCopy() *VersionAssessment {
	if in == nil {
		return nil
	}
	out := new(VersionAssessment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WinnerAssessment) DeepCopyInto(out *WinnerAssessment) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.WinnerAssessment != nil {
		in, out := &in.WinnerAssessment, &out.WinnerAssessment
		*out = new(apiv1alpha2.WinnerAssessment)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WinnerAssessment.
func (in *WinnerAssessment) DeepCopy() *WinnerAssessment {
	if in == nil {
		return nil
	}
	out := new(WinnerAssessment)
	in.DeepCopyInto(out)
	return out
}
