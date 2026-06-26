import React, { useMemo, useState } from 'react';
import {
  AppAlert,
  AppDetailSection,
  AppFilterHeader,
  AppInput,
  AppSelect,
  AppSwitch,
  AppTable,
  AppTableActionButton,
  AppTag,
  AppTooltip,
} from '../../../router-ui';

const resolveEndpointTestStatusKey = (row) =>
  (row?.last_test_status || '').toString().trim() || 'untested';

const formatTimestamp = (value) => {
  const timestamp = Number(value || 0);
  if (timestamp <= 0) {
    return '';
  }
  return new Date(timestamp * 1000).toLocaleString();
};

const resolveEndpointRuntimeState = (row) => {
  if (row?.enabled === true) {
    return 'available';
  }
  if (
    (row?.disabled_reason || '').toString().trim() ||
    Number(row?.disabled_at || 0) > 0
  ) {
    return 'paused';
  }
  if ((row?.enable_block_reason || '').toString().trim()) {
    return 'blocked';
  }
  if (resolveEndpointTestStatusKey(row) !== 'success') {
    return 'untested';
  }
  return 'disabled';
};

const ChannelDetailEndpointsTab = ({
  t,
  columnWidths,
  endpointSummaryText,
  channelEndpoints,
  channelEndpointsLoading,
  channelEndpointsError,
  buildChannelEndpointKey,
  endpointCapabilityReadonly,
  endpointMutatingKey,
  updateChannelEndpointCapability,
  channelEndpointPoliciesLoading,
  channelEndpointPolicies,
  channelEndpointPoliciesError,
  endpointPolicyReadonly,
  openEndpointPolicyEditor,
}) => {
  const buildDisableInfo = (row) => {
    const parts = [];
    const disabledBy = (row?.disabled_by || '').toString().trim();
    const disabledAt = formatTimestamp(row?.disabled_at);
    const disabledReason = (row?.disabled_reason || '').toString().trim();
    if (disabledBy) {
      parts.push(t('channel.edit.capability_disable.by', { value: disabledBy }));
    }
    if (disabledAt) {
      parts.push(t('channel.edit.capability_disable.at', { value: disabledAt }));
    }
    if (disabledReason) {
      parts.push(t('channel.edit.capability_disable.reason', { value: disabledReason }));
    }
    return parts.join('\n');
  };

  const buildRuntimeDetails = (row) => {
    const state = resolveEndpointRuntimeState(row);
    const disabledBy = (row?.disabled_by || '').toString().trim();
    const disabledAt = formatTimestamp(row?.disabled_at);
    const disabledReason = (row?.disabled_reason || '').toString().trim();
    const blockReason = (row?.enable_block_reason || '').toString().trim();
    const lastTestedAt = formatTimestamp(row?.last_tested_at);
    const details = [];
    if (state === 'available') {
      details.push(
        t(
          'channel.edit.endpoint_capabilities.status_detail.recovery_state_available',
        ),
      );
    } else if (state === 'paused') {
      details.push(
        t(
          'channel.edit.endpoint_capabilities.status_detail.recovery_state_paused',
        ),
      );
    } else if (state === 'blocked') {
      details.push(
        t(
          'channel.edit.endpoint_capabilities.status_detail.recovery_state_blocked',
        ),
      );
    } else if (state === 'untested') {
      details.push(
        t(
          'channel.edit.endpoint_capabilities.status_detail.recovery_state_untested',
        ),
      );
    } else {
      details.push(
        t(
          'channel.edit.endpoint_capabilities.status_detail.recovery_state_disabled',
        ),
      );
    }
    if (disabledReason) {
      details.push(
        t('channel.edit.endpoint_capabilities.status_detail.pause_reason', {
          value: disabledReason,
        }),
      );
    }
    if (disabledAt) {
      details.push(
        t('channel.edit.endpoint_capabilities.status_detail.paused_at', {
          value: disabledAt,
        }),
      );
    }
    if (disabledBy) {
      details.push(
        t('channel.edit.endpoint_capabilities.status_detail.paused_by', {
          value: disabledBy,
        }),
      );
    }
    if (blockReason) {
      details.push(
        t('channel.edit.endpoint_capabilities.status_detail.block_reason', {
          value: blockReason,
        }),
      );
    }
    if (lastTestedAt) {
      details.push(
        t('channel.edit.endpoint_capabilities.status_detail.last_tested_at', {
          value: lastTestedAt,
        }),
      );
    }
    return details;
  };

  const policyByKey = new Map(
    channelEndpointPolicies.map((row) => [
      buildChannelEndpointKey(row.model, row.endpoint),
      row,
    ]),
  );
  const [testStatusFilter, setTestStatusFilter] = useState('all');
  const [runtimeStateFilter, setRuntimeStateFilter] = useState('all');
  const [baseURLDrafts, setBaseURLDrafts] = useState({});

  const runtimeStateStats = useMemo(
    () =>
      channelEndpoints.reduce(
        (acc, row) => {
          const state = resolveEndpointRuntimeState(row);
          acc.total += 1;
          acc[state] = (acc[state] || 0) + 1;
          return acc;
        },
        {
          total: 0,
          available: 0,
          paused: 0,
          blocked: 0,
          untested: 0,
          disabled: 0,
        },
      ),
    [channelEndpoints],
  );

  const testStatusOptions = useMemo(
    () => [
      {
        key: 'all',
        value: 'all',
        text: t('channel.edit.endpoint_capabilities.filters.all_test_status'),
      },
      ...[
        'success',
        'failed',
        'untested',
      ].map((status) => ({
        key: status,
        value: status,
        text: t(`channel.edit.model_tester.status.${status}`),
      })),
    ],
    [t],
  );

  const runtimeStateOptions = useMemo(
    () => [
      {
        key: 'all',
        value: 'all',
        text: t('channel.edit.endpoint_capabilities.filters.all_runtime_state'),
      },
      ...['available', 'paused', 'blocked', 'untested', 'disabled'].map(
        (state) => ({
          key: state,
          value: state,
          text: t(`channel.edit.endpoint_capabilities.runtime_state.${state}`),
        }),
      ),
    ],
    [t],
  );

  const filteredRows = useMemo(
    () =>
      channelEndpoints.filter((row) => {
        const runtimeState = resolveEndpointRuntimeState(row);
        if (testStatusFilter === 'all') {
          return (
            runtimeStateFilter === 'all' ||
            runtimeState === runtimeStateFilter
          );
        }
        if (resolveEndpointTestStatusKey(row) !== testStatusFilter) {
          return false;
        }
        return (
          runtimeStateFilter === 'all' ||
          runtimeState === runtimeStateFilter
        );
      }),
    [channelEndpoints, runtimeStateFilter, testStatusFilter],
  );

  const resolveBaseURLDraft = (row, endpointKey) => {
    if (Object.prototype.hasOwnProperty.call(baseURLDrafts, endpointKey)) {
      return baseURLDrafts[endpointKey];
    }
    return row.base_url || '';
  };

  return (
    <AppDetailSection
      title={t('channel.edit.endpoint_capabilities.title')}
      titleTag='span'
      headerStart={<span className='router-toolbar-meta'>({endpointSummaryText})</span>}
    >
      <div>
        <AppAlert
          type='info'
          showIcon
          className='router-section-message'
          title={t('channel.edit.endpoint_capabilities.hint')}
        />
        <div className='router-endpoint-state-summary'>
          {['total', 'available', 'paused', 'blocked', 'untested'].map((key) => (
            <div
              className={`router-endpoint-state-card router-endpoint-state-card-${key}`}
              key={key}
            >
              <div className='router-endpoint-state-card-value'>
                {runtimeStateStats[key] || 0}
              </div>
              <div className='router-endpoint-state-card-label'>
                {t(`channel.edit.endpoint_capabilities.summary_cards.${key}`)}
              </div>
            </div>
          ))}
        </div>
        <AppFilterHeader
          className='router-toolbar-compact'
          startClassName='router-block-gap-sm'
          picker={
            <div className='router-endpoint-filter-row'>
              <AppSelect
                className='router-section-dropdown router-detail-filter-dropdown router-dropdown-min-170'
                options={runtimeStateOptions}
                value={runtimeStateFilter}
                disabled={channelEndpointsLoading || channelEndpoints.length === 0}
                placeholder={t(
                  'channel.edit.endpoint_capabilities.filters.runtime_state',
                )}
                onChange={(e, { value }) =>
                  setRuntimeStateFilter((value || 'all').toString())
                }
              />
              <AppSelect
                className='router-section-dropdown router-detail-filter-dropdown router-dropdown-min-170'
                options={testStatusOptions}
                value={testStatusFilter}
                disabled={channelEndpointsLoading || channelEndpoints.length === 0}
                placeholder={t(
                  'channel.edit.endpoint_capabilities.filters.test_status',
                )}
                onChange={(e, { value }) =>
                  setTestStatusFilter((value || 'all').toString())
                }
              />
            </div>
          }
        />
        <AppTable
          className='router-detail-table router-channel-endpoint-capability-table'
          pagination={false}
          scroll={{ x: 980 }}
          locale={{
            emptyText: channelEndpointsLoading
              ? t('channel.edit.endpoint_capabilities.loading')
              : channelEndpoints.length === 0
                ? t('channel.edit.endpoint_capabilities.empty')
                : t('channel.edit.endpoint_capabilities.filtered_empty'),
          }}
          rowKey={(row) => buildChannelEndpointKey(row.model, row.endpoint)}
          dataSource={filteredRows}
          columns={[
            {
              title: t('channel.edit.endpoint_capabilities.table.model'),
              dataIndex: 'model',
              key: 'model',
              width: columnWidths[0],
              render: (value) => (
                <span
                  className='router-cell-truncate router-monospace-value'
                  title={value}
                >
                  {value}
                </span>
              ),
            },
            {
              title: t('channel.edit.endpoint_capabilities.table.endpoint'),
              dataIndex: 'endpoint',
              key: 'endpoint',
              width: columnWidths[1],
              render: (value) => (
                <span className='router-cell-truncate' title={value}>
                  {value}
                </span>
              ),
            },
            {
              title: t('channel.edit.endpoint_capabilities.table.base_url'),
              key: 'base_url',
              width: columnWidths[2],
              render: (_, row) => {
                const endpointKey = buildChannelEndpointKey(
                  row.model,
                  row.endpoint,
                );
                const isMutating = endpointMutatingKey === endpointKey;
                const draftBaseURL = resolveBaseURLDraft(row, endpointKey);
                return (
                  <AppInput
                    className='router-section-input'
                    placeholder={t(
                      'channel.edit.endpoint_capabilities.table.base_url_placeholder',
                    )}
                    value={draftBaseURL}
                    readOnly={endpointCapabilityReadonly || isMutating}
                    onChange={(e, { value }) => {
                      setBaseURLDrafts((prev) => ({
                        ...prev,
                        [endpointKey]: (value || '').toString(),
                      }));
                    }}
                    onBlur={() => {
                      const normalizedCurrent = (row.base_url || '')
                        .toString()
                        .trim();
                      const normalizedNext = (draftBaseURL || '')
                        .toString()
                        .trim();
                      if (normalizedCurrent === normalizedNext) {
                        return;
                      }
                      updateChannelEndpointCapability(
                        {
                          ...row,
                          base_url: normalizedNext,
                        },
                        { base_url: normalizedNext, enabled: row.enabled === true },
                        { skipConfirm: true },
                      );
                    }}
                  />
                );
              },
            },
            {
              title: t('channel.edit.endpoint_capabilities.table.enabled'),
              key: 'enabled',
              width: columnWidths[3],
              align: 'center',
              render: (_, row) => {
                const endpointKey = buildChannelEndpointKey(
                  row.model,
                  row.endpoint,
                );
                const isMutating = endpointMutatingKey === endpointKey;
                const blockedReason = (row.enable_block_reason || '').trim();
                const disableInfo = buildDisableInfo(row);
                const disabled =
                  endpointCapabilityReadonly ||
                  isMutating ||
                  (!!blockedReason && row.enabled !== true);
                return (
                  <AppSwitch
                    checked={row.enabled === true}
                    disabled={disabled}
                    title={blockedReason || disableInfo || undefined}
                    onChange={(_, { checked }) =>
                      updateChannelEndpointCapability(row, {
                        enabled: checked === true,
                      })
                    }
                  />
                );
              },
            },
            {
              title: t('channel.edit.endpoint_capabilities.table.test_status'),
              key: 'test_status',
              width: columnWidths[4],
              render: (_, row) => {
                const latestStatusKey = resolveEndpointTestStatusKey(row);
                const lastTestError = (row.last_test_error || '').trim();
                const statusTag = (
                  <AppTag
                    color={
                      latestStatusKey === 'success'
                        ? 'green'
                        : latestStatusKey === 'failed'
                          ? 'red'
                          : 'grey'
                    }
                  >
                    {t(`channel.edit.model_tester.status.${latestStatusKey}`)}
                  </AppTag>
                );
                if (!lastTestError) {
                  return statusTag;
                }
                return (
                  <AppTooltip
                    title={lastTestError}
                  >
                    <span>
                      {statusTag}
                    </span>
                  </AppTooltip>
                );
              },
            },
            {
              title: t('channel.edit.endpoint_capabilities.table.runtime_status'),
              key: 'runtime_status',
              width: columnWidths[5],
              render: (_, row) => {
                const state = resolveEndpointRuntimeState(row);
                const details = buildRuntimeDetails(row);
                return (
                  <div className='router-endpoint-runtime-status'>
                    <AppTag
                      color={
                        state === 'available'
                          ? 'green'
                          : state === 'paused'
                            ? 'red'
                            : state === 'blocked'
                              ? 'orange'
                              : 'grey'
                      }
                    >
                      {t(`channel.edit.endpoint_capabilities.runtime_state.${state}`)}
                    </AppTag>
                    <div className='router-endpoint-runtime-lines'>
                      {details.map((item, index) => (
                        <div key={`${state}-${index}`}>{item}</div>
                      ))}
                    </div>
                  </div>
                );
              },
            },
            {
              title: t('channel.edit.endpoint_policies.table.policy'),
              key: 'policy',
              width: columnWidths[6],
              render: (_, row) => {
                const endpointKey = buildChannelEndpointKey(row.model, row.endpoint);
                const policyRow = policyByKey.get(endpointKey) || null;
                if (
                  channelEndpointPoliciesLoading &&
                  channelEndpointPolicies.length === 0
                ) {
                  return (
                    <span className='router-cell-truncate'>
                      {t('channel.edit.endpoint_policies.loading')}
                    </span>
                  );
                }
                return (
                  <span className='router-cell-truncate'>
                    {policyRow?.template_key || '-'}
                  </span>
                );
              },
            },
            {
              title: t('channel.edit.endpoint_policies.table.actions'),
              key: 'actions',
              width: columnWidths[7],
              render: (_, row) => (
                <AppTableActionButton
                  icon='setting'
                  title={t('channel.edit.endpoint_policies.action')}
                  disabled={endpointPolicyReadonly}
                  onClick={() => openEndpointPolicyEditor(row)}
                />
              ),
            },
          ]}
        />
        {channelEndpointsError && (
          <div className='router-error-text router-error-text-top'>
            {channelEndpointsError}
          </div>
        )}
        {channelEndpointPoliciesError && (
          <div className='router-error-text router-error-text-top'>
            {channelEndpointPoliciesError}
          </div>
        )}
      </div>
    </AppDetailSection>
  );
};

export default ChannelDetailEndpointsTab;
