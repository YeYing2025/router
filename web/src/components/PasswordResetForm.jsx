import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { API, getLogo, showError, showInfo, showSuccess } from '../helpers';
import { AppButton, AppInput, AppSection } from '../router-ui';

const PasswordResetForm = () => {
  const { t } = useTranslation();
  const [email, setEmail] = useState('');
  const [loading, setLoading] = useState(false);
  const logo = getLogo();
  const navigate = useNavigate();

  const sendResetEmail = async () => {
    if (email === '') {
      showInfo(t('messages.error.empty_email', '请输入邮箱地址')); 
      return;
    }
    setLoading(true);
    const res = await API.get(`/api/v1/public/reset_password?email=${email}`);
    const { success, message } = res.data;
    setLoading(false);
    if (success) {
      showSuccess(t('messages.success.password_reset'));
      navigate('/login');
    } else {
      showError(message);
    }
  };

  return (
    <div className='router-auth-shell'>
      <div className='router-auth-panel'>
        <AppSection className='router-auth-card'>
          <div className='router-auth-card-body'>
            <div className='router-auth-title router-auth-header'>
              <img src={logo} alt='logo' className='router-auth-logo' />
              <h2>{t('auth.reset.title')}</h2>
            </div>
            <div className='router-auth-form'>
              <AppInput
                className='router-auth-input'
                fluid
                icon='mail'
                iconPosition='left'
                placeholder={t('auth.reset.email')}
                value={email}
                onChange={(e, { value }) => setEmail(value)}
              />
              <AppButton
                className='router-auth-button router-auth-primary'
                fluid
                onClick={sendResetEmail}
                loading={loading}
              >
                {t('auth.reset.button')}
              </AppButton>
            </div>

            <div className='router-auth-message'>
              <div className='router-auth-secondary-text'>
                {t('auth.reset.remember_password')}
                <Link to='/login' className='router-auth-link'>
                  {t('auth.login.login')}
                </Link>
              </div>
            </div>
          </div>
        </AppSection>
      </div>
    </div>
  );
};

export default PasswordResetForm;
