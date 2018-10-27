## 간단한 MSA 를 만들기 위해서 만든 저장소 입니다. 
*  기본적인 boiler plate 를 만들고 있습니다.
*  서버 사이의 관계와 데이터 흐름 그리고 oauth 관련 부분을 찾는 중입니다.

### 아니.. 이거.. stop
* 작은 것부터 시작하자
1. 우선 crud 에 맞는 api 베이스를 만든다.
2. 그것을 모델로 해서 여러개의 서버를 띄운다.
 * 이때 docker 로 db 랑 여러개의 서버의 설정을 한번에 올려 본다.
3. 중간에서 인증을 할 서버를 구성한다. 
 * 여러개의 서버 사이에서 통신을 하는데 중간에 다른 공격자(?) 가 들어와서 잘못된 정보를 주거나 받으면 안되니까. 
 * 해당 서버는 자바로 만드는거에서는 eureka 를 통해서 만드는 거 같은데.. 해당 부분은.. golang 으로 만드려고 하니까 가능하면 golang 이 되는 부분을 찾아 보고 아니면 eureka를 써보도록 하자. 


 ## 처음에는 간단하게 3개 정도의 서버를 가지고 시작해 보자. 
 * 기본적으로 서버는 다 api server 이고.
 * crud 에 맞춰서 받는 method 가 다르고 
    * 이 부분은 회사에서나 쓰는 사람마다 다 다르지만.. 
    * 기왕 만드는거 가능하면 표준에 맞춰서

## 그외에 신경 써야 하는 부분 
* 보안 : auth
    * oauth 로 써 보자
* 로깅 : 로깅 범위, 로깅 로테이트. 보관 주기, 로깅에 따른 모니터링 방법 
    * logger , cfg로 일자 정하기, 모니터링? hook? from mail?
* db 부분? 
    * mysql , mongodb , redis 는 써 보자 
    * mongodb, redis 위주로 추가해 보기?
    * docker 를 이용해서 인프라를 만들때는 커널을  공유하는 부분에 넣어서 
## 추가로 알아볼 부분 
* terraform 이랑 같이 쓰는 kong api gateway 부분이 있는데 해당 부분을 참조 해서 만들어 볼까?
* GCP에서 응용할 만한 부분이 있는지 찾아보자 
* terraform 에서 응용할 만한 부분이 있는지 찾아보자
