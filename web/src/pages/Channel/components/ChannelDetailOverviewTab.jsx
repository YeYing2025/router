import React from 'react';
import {
  AppButton,
  AppDetailSection,
  AppField,
  AppFormRow,
  AppInput,
  AppSelect,
} from '../../../router-ui';

const ChannelDetailOverviewTab = ({
  t,
  inputs,
  currentProtocolOption,
  channelProtocolOptions,
  detailBasicEditing,
  detailBasicSaving,
  detailBasicEditLocked,
  detailBasicReadonly,
  channelIdentifierMaxLength,
  handleInputChange,
  cancelDetailBasicEdit,
  saveDetailBasicInfo,
  setDetailBasicEditing,
  basicConnectionFields,
  addressRoutingFields,
  protocolSelectionHintContent,
  protocolSpecificFields,
  timestamp2string,
}) => {
  return (
    <>
      <AppDetailSection
        title={t('channel.edit.detail_basic_title')}
        titleTag='span'
        headerEnd={
          detailBasicEditing ? (
            <>
              <AppButton
                type='button'
                className='router-page-button'
                onClick={cancelDetailBasicEdit}
                disabled={detailBasicSaving}
              >
                {t('channel.edit.buttons.cancel')}
              </AppButton>
              <AppButton
                type='button'
                className='router-page-button'
                color='blue'
                loading={detailBasicSaving}
                disabled={detailBasicSaving}
                onClick={saveDetailBasicInfo}
              >
                {t('channel.edit.buttons.save')}
              </AppButton>
            </>
          ) : (
            <AppButton
              type='button'
              className='router-page-button'
              color='blue'
              disabled={detailBasicEditLocked}
              onClick={() => setDetailBasicEditing(true)}
            >
              {t('common.edit')}
            </AppButton>
          )
        }
      >
        <AppFormRow>
          <AppField label={t('channel.edit.id')} readOnly>
            <AppInput
              className='router-section-input'
              value={inputs.id || '-'}
              readOnly
            />
          </AppField>
          <AppField
            label={t('channel.edit.identifier')}
            required
            readOnly={detailBasicReadonly}
          >
            <AppInput
              className='router-section-input'
              name='name'
              placeholder={t('channel.edit.identifier_placeholder')}
              onChange={handleInputChange}
              value={inputs.name}
              maxLength={channelIdentifierMaxLength}
              readOnly={detailBasicReadonly}
            />
          </AppField>
          <AppField
            label={t('channel.edit.type')}
            required={!detailBasicReadonly}
            readOnly={detailBasicReadonly}
          >
            {detailBasicReadonly ? (
              <AppInput
                className='router-section-input'
                value={currentProtocolOption?.text || inputs.protocol || '-'}
                readOnly
              />
            ) : (
              <AppSelect
                className='router-section-dropdown'
                name='protocol'
                search
                options={channelProtocolOptions}
                value={inputs.protocol}
                onChange={handleInputChange}
              />
            )}
          </AppField>
        </AppFormRow>
        {protocolSelectionHintContent}
        {basicConnectionFields}
        {addressRoutingFields}
        {protocolSpecificFields}
        <AppFormRow>
          <AppField label={t('channel.edit.created_time')} readOnly>
            <AppInput
              className='router-section-input'
              value={
                inputs.created_time ? timestamp2string(inputs.created_time) : '-'
              }
              readOnly
            />
          </AppField>
          <AppField label={t('channel.edit.updated_at')} readOnly>
            <AppInput
              className='router-section-input'
              value={
                inputs.updated_at ? timestamp2string(inputs.updated_at) : '-'
              }
              readOnly
            />
          </AppField>
        </AppFormRow>
      </AppDetailSection>
    </>
  );
};

export default ChannelDetailOverviewTab;
