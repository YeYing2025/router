import React from 'react';
import {
  AppAlert,
  AppButton,
  AppDetailSection,
  AppEmpty,
  AppInput,
  AppPagination,
  AppSelect,
  AppTable,
  AppTag,
} from '../../../router-ui';

const ChannelDetailModelsTab = ({
  t,
  columnWidths,
  modelSectionMetaText,
  detailModelFilter,
  setDetailModelFilter,
  detailModelsEditing,
  modelSearchKeyword,
  setModelSearchKeyword,
  fetchModelsLoading,
  activeRefreshModelsTask,
  detailModelMutating,
  handleFetchModels,
  searchedModelConfigs,
  visibleModelConfigs,
  renderedModelConfigs,
  getComplexPricingDetailsForModel,
  openComplexPricingModal,
  detailModelsEditLocked,
  providerCatalogLoading,
  toggleModelSelection,
  canSelectChannelModel,
  detailCurrentPageAllSelected,
  detailCurrentPagePartiallySelected,
  detailCurrentPageSelectableCount,
  toggleDetailCurrentPageSelections,
  normalizeChannelModelType,
  startDetailModelEdit,
  detailModelTotalPages,
  detailModelPage,
  setDetailModelPage,
  modelsSyncError,
}) => {
  const tableRowSelection = {
    columnWidth: columnWidths[0],
    selectedRowKeys: renderedModelConfigs
      .filter((row) => row?.selected)
      .map((row) => `${row.upstream_model}-${row.model}`),
    getTitleCheckboxProps: () => ({
      checked: detailCurrentPageAllSelected,
      indeterminate: detailCurrentPagePartiallySelected,
      disabled:
        detailModelsEditing ||
        detailModelMutating ||
        providerCatalogLoading ||
        detailCurrentPageSelectableCount === 0,
    }),
    getCheckboxProps: (row) => {
      const canSelect = canSelectChannelModel(row);
      const isUnavailable = !canSelect && !row.selected;
      const disabledReason = isUnavailable
        ? t('channel.edit.model_selector.selection_disabled_unassigned')
        : '';
      return {
        className: isUnavailable ? 'router-model-toggle-disabled' : undefined,
        disabled:
          detailModelMutating ||
          detailModelsEditing ||
          providerCatalogLoading ||
          isUnavailable,
        title: disabledReason || undefined,
      };
    },
    renderCell: (_, row, __, originNode) => {
      const canSelect = canSelectChannelModel(row);
      const isUnavailable = !canSelect && !row.selected;
      const disabledReason = isUnavailable
        ? t('channel.edit.model_selector.selection_disabled_unassigned')
        : '';
      return (
        <span
          className={[
            'router-inline-block',
            'router-model-toggle-wrap',
            isUnavailable ? 'router-model-toggle-wrap-disabled' : '',
          ]
            .filter(Boolean)
            .join(' ')}
          title={disabledReason || undefined}
          aria-label={disabledReason || undefined}
        >
          {originNode}
        </span>
      );
    },
    onSelect: (record, selected) => {
      toggleModelSelection(record.upstream_model, selected);
    },
    onSelectAll: (selected) => {
      toggleDetailCurrentPageSelections(selected);
    },
  };

  return (
    <AppDetailSection
      title={t('channel.edit.detail_models_title')}
      titleTag='span'
      headerStart={<span className='router-toolbar-meta'>({modelSectionMetaText})</span>}
      headerEnd={
        <>
          <AppSelect
            className='router-section-dropdown router-dropdown-min-170 router-detail-filter-dropdown'
            disabled={detailModelsEditing}
            options={[
              {
                key: 'all',
                value: 'all',
                text: t('channel.edit.model_selector.filters.all'),
              },
              {
                key: 'enabled',
                value: 'enabled',
                text: t('channel.edit.model_selector.filters.enabled'),
              },
              {
                key: 'disabled',
                value: 'disabled',
                text: t('channel.edit.model_selector.filters.disabled'),
              },
            ]}
            value={detailModelFilter}
            onChange={(e, { value }) =>
              setDetailModelFilter((value || 'all').toString())
            }
          />
          <AppInput
            className='router-section-input router-search-form-sm'
            icon='search'
            iconPosition='left'
            disabled={detailModelsEditing}
            placeholder={t('channel.edit.model_selector.search_placeholder')}
            value={modelSearchKeyword}
            onChange={(e, { value }) => setModelSearchKeyword(value || '')}
          />
          <AppButton
            type='button'
            className='router-page-button'
            loading={fetchModelsLoading || !!activeRefreshModelsTask}
            disabled={
              detailModelsEditing ||
              fetchModelsLoading ||
              !!activeRefreshModelsTask ||
              detailModelMutating
            }
            onClick={() => handleFetchModels({ silent: false })}
          >
            {t('channel.edit.buttons.sync_models')}
          </AppButton>
        </>
      }
    >
      <div>
        <AppAlert
          type='info'
          showIcon
          className='router-section-message'
          title={t('channel.edit.model_selector.enable_hint')}
        />
        <AppTable
          className='router-detail-table router-channel-detail-model-table'
          pagination={false}
          scroll={{ x: 1120 }}
          rowSelection={tableRowSelection}
          locale={{
            emptyText: (
              <AppEmpty>
                {modelSearchKeyword.trim() !== ''
                  ? t('channel.edit.model_selector.empty_search')
                  : visibleModelConfigs.length > 0
                    ? t('channel.edit.model_selector.empty_filtered')
                    : t('channel.edit.model_selector.empty')}
              </AppEmpty>
            ),
          }}
          rowKey={(row) => `${row.upstream_model}-${row.model}`}
          dataSource={searchedModelConfigs.length === 0 ? [] : renderedModelConfigs}
          columns={[
            {
              title: t('channel.edit.model_selector.table.name'),
              dataIndex: 'upstream_model',
              key: 'upstream_model',
              width: columnWidths[1],
              render: (value, row) => (
                <div className='router-cell-truncate' title={value}>
                  <span className='router-nowrap'>{value}</span>
                  {row.inactive && (
                    <AppTag color='grey' className='router-tag'>
                      {t('channel.edit.model_selector.inactive')}
                    </AppTag>
                  )}
                </div>
              ),
            },
            {
              title: t('channel.edit.model_selector.table.type'),
              key: 'type',
              width: columnWidths[2],
              render: (_, row) =>
                t(`channel.model_types.${normalizeChannelModelType(row.type)}`),
            },
            {
              title: t('channel.edit.model_selector.table.alias'),
              dataIndex: 'model',
              key: 'model',
              width: columnWidths[3],
              render: (value) => (
                <span className='router-cell-truncate' title={value}>
                  {value}
                </span>
              ),
            },
            {
              title: t('channel.edit.model_selector.table.price_unit'),
              dataIndex: 'price_unit',
              key: 'price_unit',
              width: columnWidths[4],
              render: (value) => <span className='router-nowrap'>{value}</span>,
            },
            {
              title: t('channel.edit.model_selector.table.input_price'),
              key: 'input_price',
              width: columnWidths[5],
              render: (_, row) => {
                const complexPricingDetails = getComplexPricingDetailsForModel(row);
                const hasComplexInputPricing = complexPricingDetails.some((detail) =>
                  (detail.price_components || []).some(
                    (component) => Number(component.input_price || 0) > 0,
                  ),
                );
                if (hasComplexInputPricing) {
                  return (
                    <AppButton
                      type='button'
                      className='router-inline-button'
                      onClick={() => openComplexPricingModal(row)}
                    >
                      {t('channel.edit.model_selector.pricing_detail_button')}
                    </AppButton>
                  );
                }
                return <span className='router-nowrap'>{row.input_price ?? '-'}</span>;
              },
            },
            {
              title: t('channel.edit.model_selector.table.output_price'),
              key: 'output_price',
              width: columnWidths[6],
              render: (_, row) => {
                const complexPricingDetails = getComplexPricingDetailsForModel(row);
                const hasComplexOutputPricing = complexPricingDetails.some((detail) =>
                  (detail.price_components || []).some(
                    (component) => Number(component.output_price || 0) > 0,
                  ),
                );
                if (hasComplexOutputPricing) {
                  return (
                    <AppButton
                      type='button'
                      className='router-inline-button'
                      onClick={() => openComplexPricingModal(row)}
                    >
                      {t('channel.edit.model_selector.pricing_detail_button')}
                    </AppButton>
                  );
                }
                return <span className='router-nowrap'>{row.output_price ?? '-'}</span>;
              },
            },
            {
              title: t('channel.table.actions'),
              key: 'actions',
              width: columnWidths[7],
              render: (_, row) => {
                const rowEditDisabled =
                  detailModelsEditLocked || detailModelMutating || detailModelsEditing;
                const rowActionBlocked = !canSelectChannelModel(row) && !row.selected;
                const rowActionDisabled = rowEditDisabled || rowActionBlocked;
                const rowActionDisabledReason = rowActionBlocked
                  ? t('channel.edit.model_selector.selection_disabled_unassigned')
                  : '';
                return (
                  <div className='router-inline-actions'>
                    <AppButton
                      type='button'
                      className='router-inline-button'
                      disabled={rowActionDisabled}
                      title={rowActionDisabledReason || undefined}
                      onClick={() => startDetailModelEdit(row.upstream_model)}
                    >
                      {t('common.edit')}
                    </AppButton>
                  </div>
                );
              },
            },
          ]}
        />
        {detailModelTotalPages > 1 && (
          <div className='router-pagination-wrap'>
            <AppPagination
              className='router-section-pagination'
              activePage={detailModelPage}
              totalPages={detailModelTotalPages}
              onPageChange={(e, { activePage }) =>
                setDetailModelPage(Number(activePage) || 1)
              }
            />
          </div>
        )}
        {modelsSyncError && (
          <div className='router-error-text router-error-text-top'>
            {modelsSyncError}
          </div>
        )}
      </div>
    </AppDetailSection>
  );
};

export default ChannelDetailModelsTab;
