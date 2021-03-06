package consumptionapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/consumption/mgmt/2018-08-31/consumption"
	"github.com/Azure/go-autorest/autorest"
)

// UsageDetailsClientAPI contains the set of methods on the UsageDetailsClient type.
type UsageDetailsClientAPI interface {
	List(ctx context.Context, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByBillingAccount(ctx context.Context, billingAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListByDepartment(ctx context.Context, departmentID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByEnrollmentAccount(ctx context.Context, enrollmentAccountID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListByManagementGroup(ctx context.Context, managementGroupID string, expand string, filter string, skiptoken string, top *int32, apply string) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByBillingAccount(ctx context.Context, billingAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByDepartment(ctx context.Context, departmentID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByEnrollmentAccount(ctx context.Context, enrollmentAccountID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
	ListForBillingPeriodByManagementGroup(ctx context.Context, managementGroupID string, billingPeriodName string, expand string, filter string, apply string, skiptoken string, top *int32) (result consumption.UsageDetailsListResultPage, err error)
}

var _ UsageDetailsClientAPI = (*consumption.UsageDetailsClient)(nil)

// MarketplacesClientAPI contains the set of methods on the MarketplacesClient type.
type MarketplacesClientAPI interface {
	List(ctx context.Context, filter string, top *int32, skiptoken string) (result consumption.MarketplacesListResultPage, err error)
	ListByBillingAccount(ctx context.Context, billingAccountID string, filter string, top *int32, skiptoken string) (result consumption.MarketplacesListResultPage, err error)
	ListByBillingPeriod(ctx context.Context, billingPeriodName string, filter string, top *int32, skiptoken string) (result consumption.MarketplacesListResultPage, err error)
	ListByDepartment(ctx context.Context, departmentID string, filter string, top *int32, skiptoken string) (result consumption.MarketplacesListResultPage, err error)
	ListByEnrollmentAccount(ctx context.Context, enrollmentAccountID string, filter string, top *int32, skiptoken string) (result consumption.MarketplacesListResultPage, err error)
	ListForBillingPeriodByBillingAccount(ctx context.Context, billingAccountID string, billingPeriodName string, filter string, top *int32, skiptoken string) (result consumption.MarketplacesListResultPage, err error)
	ListForBillingPeriodByDepartment(ctx context.Context, departmentID string, billingPeriodName string, filter string, top *int32, skiptoken string) (result consumption.MarketplacesListResultPage, err error)
	ListForBillingPeriodByEnrollmentAccount(ctx context.Context, enrollmentAccountID string, billingPeriodName string, filter string, top *int32, skiptoken string) (result consumption.MarketplacesListResultPage, err error)
}

var _ MarketplacesClientAPI = (*consumption.MarketplacesClient)(nil)

// BalancesClientAPI contains the set of methods on the BalancesClient type.
type BalancesClientAPI interface {
	GetByBillingAccount(ctx context.Context, billingAccountID string) (result consumption.Balance, err error)
	GetForBillingPeriodByBillingAccount(ctx context.Context, billingAccountID string, billingPeriodName string) (result consumption.Balance, err error)
}

var _ BalancesClientAPI = (*consumption.BalancesClient)(nil)

// ReservationsSummariesClientAPI contains the set of methods on the ReservationsSummariesClient type.
type ReservationsSummariesClientAPI interface {
	ListByReservationOrder(ctx context.Context, reservationOrderID string, grain consumption.Datagrain, filter string) (result consumption.ReservationSummariesListResultPage, err error)
	ListByReservationOrderAndReservation(ctx context.Context, reservationOrderID string, reservationID string, grain consumption.Datagrain, filter string) (result consumption.ReservationSummariesListResultPage, err error)
}

var _ ReservationsSummariesClientAPI = (*consumption.ReservationsSummariesClient)(nil)

// ReservationsDetailsClientAPI contains the set of methods on the ReservationsDetailsClient type.
type ReservationsDetailsClientAPI interface {
	ListByReservationOrder(ctx context.Context, reservationOrderID string, filter string) (result consumption.ReservationDetailsListResultPage, err error)
	ListByReservationOrderAndReservation(ctx context.Context, reservationOrderID string, reservationID string, filter string) (result consumption.ReservationDetailsListResultPage, err error)
}

var _ ReservationsDetailsClientAPI = (*consumption.ReservationsDetailsClient)(nil)

// ReservationRecommendationsClientAPI contains the set of methods on the ReservationRecommendationsClient type.
type ReservationRecommendationsClientAPI interface {
	List(ctx context.Context, filter string) (result consumption.ReservationRecommendationsListResultPage, err error)
}

var _ ReservationRecommendationsClientAPI = (*consumption.ReservationRecommendationsClient)(nil)

// BudgetsClientAPI contains the set of methods on the BudgetsClient type.
type BudgetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, budgetName string, parameters consumption.Budget) (result consumption.Budget, err error)
	CreateOrUpdateByResourceGroupName(ctx context.Context, resourceGroupName string, budgetName string, parameters consumption.Budget) (result consumption.Budget, err error)
	Delete(ctx context.Context, budgetName string) (result autorest.Response, err error)
	DeleteByResourceGroupName(ctx context.Context, resourceGroupName string, budgetName string) (result autorest.Response, err error)
	Get(ctx context.Context, budgetName string) (result consumption.Budget, err error)
	GetByResourceGroupName(ctx context.Context, resourceGroupName string, budgetName string) (result consumption.Budget, err error)
	List(ctx context.Context) (result consumption.BudgetsListResultPage, err error)
	ListByResourceGroupName(ctx context.Context, resourceGroupName string) (result consumption.BudgetsListResultPage, err error)
}

var _ BudgetsClientAPI = (*consumption.BudgetsClient)(nil)

// PriceSheetClientAPI contains the set of methods on the PriceSheetClient type.
type PriceSheetClientAPI interface {
	Get(ctx context.Context, expand string, skiptoken string, top *int32) (result consumption.PriceSheetResult, err error)
	GetByBillingPeriod(ctx context.Context, billingPeriodName string, expand string, skiptoken string, top *int32) (result consumption.PriceSheetResult, err error)
}

var _ PriceSheetClientAPI = (*consumption.PriceSheetClient)(nil)

// TagsClientAPI contains the set of methods on the TagsClient type.
type TagsClientAPI interface {
	Get(ctx context.Context, billingAccountID string) (result consumption.TagsResult, err error)
}

var _ TagsClientAPI = (*consumption.TagsClient)(nil)

// ForecastsClientAPI contains the set of methods on the ForecastsClient type.
type ForecastsClientAPI interface {
	List(ctx context.Context, filter string) (result consumption.ForecastsListResult, err error)
}

var _ ForecastsClientAPI = (*consumption.ForecastsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result consumption.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*consumption.OperationsClient)(nil)

// AggregatedCostClientAPI contains the set of methods on the AggregatedCostClient type.
type AggregatedCostClientAPI interface {
	GetByManagementGroup(ctx context.Context, managementGroupID string, filter string) (result consumption.ManagementGroupAggregatedCostResult, err error)
	GetForBillingPeriodByManagementGroup(ctx context.Context, managementGroupID string, billingPeriodName string) (result consumption.ManagementGroupAggregatedCostResult, err error)
}

var _ AggregatedCostClientAPI = (*consumption.AggregatedCostClient)(nil)

// ChargesClientAPI contains the set of methods on the ChargesClient type.
type ChargesClientAPI interface {
	ListByDepartment(ctx context.Context, billingAccountID string, departmentID string, filter string) (result consumption.ChargesListResult, err error)
	ListByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, filter string) (result consumption.ChargesListResult, err error)
	ListForBillingPeriodByDepartment(ctx context.Context, billingAccountID string, departmentID string, billingPeriodName string, filter string) (result consumption.ChargeSummary, err error)
	ListForBillingPeriodByEnrollmentAccount(ctx context.Context, billingAccountID string, enrollmentAccountID string, billingPeriodName string, filter string) (result consumption.ChargeSummary, err error)
}

var _ ChargesClientAPI = (*consumption.ChargesClient)(nil)
