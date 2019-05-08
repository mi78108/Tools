require 'mysql2'
class Mysql_tool
  def initialize(host_addr='127.0.0.1',user_name='root',user_passwd='mysql',database='test',max_connect=27)
		@max_connect=max_connect
		@host=host_addr
		@username=user_name
		@password=user_passwd
		@database=database
		@connect_pool=Array.new
		@connect_size=0
		@mutex=Mutex.new
	end

	def get_new_connect
		conn = Mysql2::Client.new(
			:host     => @host, 	#'127.0.0.1', # 主机
			:username => @username, # 用户名
			:password => @password, # 密码
			:database => @database, # 数据库
			:encoding => 'utf8'     # 编码
		)
		@connect_size+=1 if  !conn.nil?		
		return conn
	end

	def get_pool_connect
		conn=nil
		@mutex.synchronize{
			begin
				conn=@connect_pool.pop
				break if !conn.nil?
				if((@connect_pool.size < @max_connect) and @connect_pool.size == 0)
					conn=get_new_connect
				end
			end while conn==nil
		}
		return conn
	end

	def release_connect(conn)
		@connect_pool.push(conn)
	end

	def pool_query(sql_string)
		db=get_pool_connect
		r=db.query(sql_string)
		release_connect(db)
		return r
	end
end
