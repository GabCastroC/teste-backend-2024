module Services
  module Api 
    module V1
      module Products
        class Upsert
          attr_accessor :params, :request

          def initialize(params, request)
            @params = params
            @request = request
          end

          def execute
            # raise ActiveRecord::RecordNotFound, "Product Not Found" if product.blank?
            is_create_event = product.new_record?
            
            ActiveRecord::Base.transaction do
              product.id        ||= params[:id]           if params[:id].present?
              product.name        = params[:name]         if params[:name].present?
              product.amount       = params[:amount]        if params[:amount].present?
              product.brand       = params[:brand]        if params[:brand].present?
              product.price       = params[:price]        if params[:price].present?
              product.description = params[:description]  if params[:description].present?

              product.save!

            end
            
            if params[:is_api]
              event_type = is_create_event ? 'create' : 'update'
              payload = product.as_json.merge(event_type: event_type)
              Karafka.producer.produce_sync(topic: 'rails-to-go', payload: payload.to_json)
            end
          end

          private
          def product
            @product ||= Product.find_by(id: params[:id]).presence || Product.new
          end
        end
      end
    end
  end
end