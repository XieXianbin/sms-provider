# �������api�ӿ�

==================

*ɳ�价���������� ��˴˹�������*

���÷���

���ð�

```go
import(
  "alidayu"
  )
```

��ʼ������


```go
// ��ʼ��������

alidayu.Appkey = "xxxxxxx"

alidayu.AppSecret = "xxxxxxxxxxxxxxxxxxxxxxxxx"

alidayu.IsDebug = false
```

## ���Žӿ�

* SendOnce ���ŵ������͵Ľӿ�

>// moblie-�ֻ�����

>// signname-����ǩ��

>// templatecode-����ģ��

>// param-�������

>// ����Result��ʽ����ȷ������ɹ�ʧ�ܵĽṹ�����ã����Result.SuccessΪtrue���ȡResult.ResultError����

```go
alidayu.SendOnce("136xxxxxxx8","����",��SM_777777��,"{'code':'666666'}")
```

##  

* SendBatch ���ŵ������͵Ľӿ�

>// moblie-�ֻ����� ','����

>// signname-����ǩ��

>// templatecode-����ģ��

>// param-�������

>// ����Result��ʽ����ȷ������ɹ�ʧ�ܵĽṹ�����ã����Result.SuccessΪtrue���ȡResult.ResultError����

```go
alidayu.SendBatch("136xxxxxxx8,136xxxxxxx3","����","SM_777777","{'code':'666666'}")
```

## �ı�ת�����ӿ�

* SendLecall �ı�ת�����Ľӿ�

>// moblie-�ֻ�����

>// templatecode-����ģ��

>// param-�������

>// ����Result��ʽ����ȷ������ɹ�ʧ�ܵĽṹ�����ã����Result.SuccessΪtrue���ȡResult.ResultError����

```go
alidayu.SendLecall("136xxxxxxx8","TL_777777","{'code':��666666��}")
```


## ���ظ�ʽ

���ظ�ʽĬ��Ϊmodel.go�µ�Result��ʽ

������ʱ��������Ϊ�յ��ӽṹ

����Result.Success==trueʱ

���õ�Result.resultError�ᱨ��

## ����method��ķ���

����ת���ֵķ�ʽ���Ĭ�Ͻ����ֿɶ���

�������Ҫ������֤�� ����Ҫ�������ö��Ÿ���

��ʱ���Ե���SpritCode����

���Ὣ����666666ת��Ϊ6,6,6,6,6,6

������������

##  

������Դ����ǲ�������ڿ���ʱ�����������Դ��Ŀ��
