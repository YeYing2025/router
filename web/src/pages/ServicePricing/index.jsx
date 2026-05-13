import React from 'react';
import { useTranslation } from 'react-i18next';
import BalanceTopUpPage from '../TopUp/BalanceTopUpPage';
import PackagePurchasePage from '../TopUp/PackagePurchasePage';
import TopUpWorkspaceProvider from '../TopUp/provider.jsx';
import { AppFilterHeader, AppSection } from '../../router-ui';

const ServicePricing = () => {
  const { t } = useTranslation();

  return (
    <TopUpWorkspaceProvider>
      <div className='dashboard-container'>
        <AppSection>
          <div className='router-service-pricing-hero'>
            <AppFilterHeader
              className='router-block-gap-md'
              title={t('topup.pricing.page_title')}
              titleClassName='router-ui-section-title router-service-pricing-title'
              meta={t('topup.pricing.subtitle')}
            />
            <div className='router-service-pricing-mode-grid'>
              <div className='router-service-pricing-mode-card router-service-pricing-mode-card-package'>
                <div className='router-service-pricing-mode-title router-service-pricing-mode-title-package'>
                  {t('topup.pricing.package_mode_title')}
                </div>
                <div className='router-service-pricing-mode-body'>
                  {t('topup.pricing.package_hint')}
                </div>
              </div>
              <div className='router-service-pricing-mode-card router-service-pricing-mode-card-balance'>
                <div className='router-service-pricing-mode-title router-service-pricing-mode-title-balance'>
                  {t('topup.pricing.balance_mode_title')}
                </div>
                <div className='router-service-pricing-mode-body'>
                  {t('topup.pricing.balance_hint')}
                </div>
              </div>
            </div>
          </div>
        </AppSection>

        <PackagePurchasePage />
        <BalanceTopUpPage showCurrentBalance={false} />
      </div>
    </TopUpWorkspaceProvider>
  );
};

export default ServicePricing;
