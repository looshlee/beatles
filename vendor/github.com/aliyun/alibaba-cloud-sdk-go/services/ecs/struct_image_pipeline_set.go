package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ImagePipelineSet is a nested struct in ecs response
type ImagePipelineSet struct {
	CreationTime            string                       `json:"CreationTime" xml:"CreationTime"`
	ImagePipelineId         string                       `json:"ImagePipelineId" xml:"ImagePipelineId"`
	Name                    string                       `json:"Name" xml:"Name"`
	Description             string                       `json:"Description" xml:"Description"`
	BaseImageType           string                       `json:"BaseImageType" xml:"BaseImageType"`
	BaseImage               string                       `json:"BaseImage" xml:"BaseImage"`
	ImageName               string                       `json:"ImageName" xml:"ImageName"`
	VSwitchId               string                       `json:"VSwitchId" xml:"VSwitchId"`
	InstanceType            string                       `json:"InstanceType" xml:"InstanceType"`
	InternetMaxBandwidthOut int                          `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	SystemDiskSize          int                          `json:"SystemDiskSize" xml:"SystemDiskSize"`
	DeleteInstanceOnFailure bool                         `json:"DeleteInstanceOnFailure" xml:"DeleteInstanceOnFailure"`
	BuildContent            string                       `json:"BuildContent" xml:"BuildContent"`
	ResourceGroupId         string                       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	AddAccounts             AddAccounts                  `json:"AddAccounts" xml:"AddAccounts"`
	ToRegionIds             ToRegionIds                  `json:"ToRegionIds" xml:"ToRegionIds"`
	Tags                    TagsInDescribeImagePipelines `json:"Tags" xml:"Tags"`
}
