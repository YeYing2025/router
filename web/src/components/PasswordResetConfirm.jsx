import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { API, copy, getLogo, showError, showNotice } from '../helpers';
import { useSearchParams } from 'react-router-dom';
import { AppAlert, AppButton, AppInput, AppSection } from '../router-ui';

const PasswordResetConfirm = () => {
  const { t } = useTranslation();
  const [inputs, setInputs] = useState({
    email: '',
    token: '',
  });
  const { email, token } = inputs;
  const [loading, setLoading] = useState(false);
  const [disableButton, setDisableButton] = useState(false);
  const [newPassword, setNewPassword] = useState('');
  const logo = getLogo();

  const [countdown, setCountdown] = useState(30);

  const [searchParams] = useSearchParams();
  useEffect(() => {
    let token = searchParams.get('token');
    let email = searchParams.get('email');
    setInputs({
      token,
      email,
    });
  }, [searchParams]);

  useEffect(() => {
    let countdownInterval = null;
    if (disableButton && countdown > 0) {
      countdownInterval = setInterval(() => {
        setCountdown(countdown - 1);
      }, 1000);
    } else if (countdown === 0) {
      setDisableButton(false);
      setCountdown(30);
    }
    return () => clearInterval(countdownInterval);
  }, [disableButton, countdown]);

  async function handleSubmit(e) {
    setDisableButton(true);
    if (!email) return;
    setLoading(true);
    const res = await API.post(`/api/v1/public/user/reset`, {
      email,
      token,
    });
    const { success, message } = res.data;
    if (success) {
      let password = res.data.data;
      setNewPassword(password);
      await copy(password);
      showNotice(t('messages.notice.password_copied', { password }));
    } else {
      showError(message);
    }
    setLoading(false);
  }

  return (
    <div className='router-auth-shell'>
      <div className='router-auth-panel'>
        <AppSection className='router-auth-card'>
          <div className='router-auth-card-body'>
            <div className='router-auth-title router-auth-header'>
              <img src={logo} alt='logo' className='router-auth-logo' />
              <h2>{t('auth.reset.confirm.title')}</h2>
            </div>
            <div className='router-auth-form'>
              <AppInput
                className='router-auth-input'
                fluid
                icon='mail'
                iconPosition='left'
                placeholder={t('auth.reset.email')}
                name='email'
                value={email}
                readOnly
              />
              {newPassword && (
                <AppInput
                  className='router-auth-input router-auth-input-clickable'
                  fluid
                  icon='lock'
                  iconPosition='left'
                  placeholder={t('auth.reset.confirm.new_password')}
                  name='newPassword'
                  value={newPassword}
                  readOnly
                  onClick={(e) => {
                    e.target.select();
                    navigator.clipboard.writeText(newPassword);
                    showNotice(t('auth.reset.confirm.notice'));
                  }}
                />
              )}
              <AppButton
                className='router-auth-button router-auth-primary'
                fluid
                onClick={handleSubmit}
                loading={loading}
                disabled={disableButton}
              >
                {disableButton
                  ? t('auth.reset.confirm.button_disabled')
                  : t('auth.reset.confirm.button')}
              </AppButton>
            </div>
            {newPassword && (
              <AppAlert
                type='info'
                showIcon
                className='router-auth-message'
                title={t('auth.reset.confirm.notice')}
              />
            )}
          </div>
        </AppSection>
      </div>
    </div>
  );
};

export default PasswordResetConfirm;
