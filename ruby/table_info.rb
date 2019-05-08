require 'mysql2'

Db=Mysql2::Client.new(:host=>'192.168.2.8',:username=>'jklz',:password=>'sQWeBvaqT6B3',:database=>'his')

tables=Db.query(%Q{show tables;}).to_a.map{|t| t['Tables_in_his']}

f=File.open('./tables.info','w')
tables.each do |t|
  t_r=Db.query(%Q{select TABLE_NAME,TABLE_COMMENT from information_schema.TABLES where TABLE_SCHEMA='his' and TABLE_NAME='#{t}';}).to_a.first
  t_c=Db.query(%Q{select COLUMN_NAME,COLUMN_COMMENT from information_schema.COLUMNS where TABLE_SCHEMA='his' and TABLE_NAME='#{t}';}).to_a
  f.puts("----#{t}------")
  f.puts("T: #{t_r['TABLE_NAME']} ---> #{t_r['TABLE_COMMENT']}")
  t_c.each do |c|
    f.puts("C: #{c['COLUMN_NAME']} ---> #{c['COLUMN_COMMENT']}")
  end
  f.puts("--------------")
  f.write("\n")
end
f.close
